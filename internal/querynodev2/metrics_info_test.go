// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package querynodev2

import (
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	"github.com/milvus-io/milvus/internal/distributed/streaming"
	"github.com/milvus-io/milvus/internal/json"
	"github.com/milvus-io/milvus/internal/mocks/distributed/mock_streaming"
	"github.com/milvus-io/milvus/internal/querynodev2/delegator"
	"github.com/milvus-io/milvus/internal/querynodev2/pipeline"
	"github.com/milvus-io/milvus/internal/querynodev2/segments"
	"github.com/milvus-io/milvus/pkg/v2/mq/msgdispatcher"
	"github.com/milvus-io/milvus/pkg/v2/proto/querypb"
	"github.com/milvus-io/milvus/pkg/v2/streaming/util/types"
	"github.com/milvus-io/milvus/pkg/v2/util/metricsinfo"
	"github.com/milvus-io/milvus/pkg/v2/util/paramtable"
	"github.com/milvus-io/milvus/pkg/v2/util/tsoutil"
	"github.com/milvus-io/milvus/pkg/v2/util/typeutil"
)

func TestGetPipelineJSON(t *testing.T) {
	paramtable.Init()

	ch := "ch"
	delegators := typeutil.NewConcurrentMap[string, delegator.ShardDelegator]()
	d := delegator.NewMockShardDelegator(t)
	d.EXPECT().GetTSafe().Return(0)
	delegators.Insert(ch, d)
	msgDispatcher := msgdispatcher.NewMockClient(t)

	collectionManager := segments.NewMockCollectionManager(t)
	segmentManager := segments.NewMockSegmentManager(t)
	collectionManager.EXPECT().Get(mock.Anything).Return(segments.NewTestCollection(1, querypb.LoadType_UnKnownType, &schemapb.CollectionSchema{}))
	manager := &segments.Manager{
		Collection: collectionManager,
		Segment:    segmentManager,
	}

	pipelineManager := pipeline.NewManager(manager, msgDispatcher, delegators)

	_, err := pipelineManager.Add(1, ch)
	assert.NoError(t, err)
	assert.Equal(t, 1, pipelineManager.Num())

	stats := pipelineManager.GetChannelStats(0)
	expectedStats := []*metricsinfo.Channel{
		{
			Name:           ch,
			WatchState:     "Healthy",
			LatestTimeTick: tsoutil.PhysicalTimeFormat(0),
			NodeID:         paramtable.GetNodeID(),
			CollectionID:   1,
		},
	}
	assert.Equal(t, expectedStats, stats)

	JSONStr := getChannelJSON(&QueryNode{pipelineManager: pipelineManager}, 0)
	assert.NotEmpty(t, JSONStr)

	var actualStats []*metricsinfo.Channel
	err = json.Unmarshal([]byte(JSONStr), &actualStats)
	assert.NoError(t, err)
	assert.Equal(t, expectedStats, actualStats)
}

func TestGetSegmentJSON(t *testing.T) {
	segment := segments.NewMockSegment(t)
	segment.EXPECT().ID().Return(int64(1))
	segment.EXPECT().Collection().Return(int64(1001))
	segment.EXPECT().Partition().Return(int64(2001))
	segment.EXPECT().MemSize().Return(int64(1024))
	segment.EXPECT().HasRawData(mock.Anything).Return(true)
	segment.EXPECT().Indexes().Return([]*segments.IndexedFieldInfo{
		{
			IndexInfo: &querypb.FieldIndexInfo{
				FieldID:   1,
				IndexID:   101,
				IndexSize: 512,
				BuildID:   10001,
			},
			IsLoaded: true,
		},
	})
	segment.EXPECT().Type().Return(segments.SegmentTypeGrowing)
	segment.EXPECT().ResourceGroup().Return("default")
	segment.EXPECT().InsertCount().Return(int64(100))

	node := &QueryNode{}
	mockedSegmentManager := segments.NewMockSegmentManager(t)
	mockedSegmentManager.EXPECT().GetBy().Return([]segments.Segment{segment})
	node.manager = &segments.Manager{Segment: mockedSegmentManager}

	jsonStr := getSegmentJSON(node, 0)
	assert.NotEmpty(t, jsonStr)

	var segments []*metricsinfo.Segment
	err := json.Unmarshal([]byte(jsonStr), &segments)
	assert.NoError(t, err)
	assert.NotNil(t, segments)
	assert.Equal(t, 1, len(segments))
	assert.Equal(t, int64(1), segments[0].SegmentID)
	assert.Equal(t, int64(1001), segments[0].CollectionID)
	assert.Equal(t, int64(2001), segments[0].PartitionID)
	assert.Equal(t, int64(1024), segments[0].MemSize)
	assert.Equal(t, 1, len(segments[0].IndexedFields))
	assert.Equal(t, int64(1), segments[0].IndexedFields[0].IndexFieldID)
	assert.Equal(t, int64(101), segments[0].IndexedFields[0].IndexID)
	assert.Equal(t, int64(512), segments[0].IndexedFields[0].IndexSize)
	assert.Equal(t, int64(10001), segments[0].IndexedFields[0].BuildID)
	assert.True(t, segments[0].IndexedFields[0].IsLoaded)
	assert.Equal(t, "Growing", segments[0].State)
	assert.Equal(t, "default", segments[0].ResourceGroup)
	assert.Equal(t, int64(100), segments[0].LoadedInsertRowCount)
}

func TestStreamingQuotaMetrics(t *testing.T) {
	paramtable.Init()

	wal := mock_streaming.NewMockWALAccesser(t)
	local := mock_streaming.NewMockLocal(t)
	now := time.Now()
	local.EXPECT().GetMetricsIfLocal(mock.Anything).Return(&types.StreamingNodeMetrics{
		WALMetrics: map[types.ChannelID]types.WALMetrics{
			{Name: "ch1"}: types.RWWALMetrics{
				ChannelInfo: types.PChannelInfo{
					Name: "ch1",
				},
				MVCCTimeTick:     tsoutil.ComposeTSByTime(now, 0),
				RecoveryTimeTick: tsoutil.ComposeTSByTime(now.Add(-time.Second), 0),
			},
			{Name: "ch2"}: types.ROWALMetrics{},
		},
	}, nil)
	wal.EXPECT().Local().Return(local)
	streaming.SetWALForTest(wal)
	defer streaming.RecoverWALForTest()

	m := getStreamingQuotaMetrics()
	assert.Len(t, m.WALs, 1)
	assert.Equal(t, "ch1", m.WALs[0].Channel.Name)
	assert.Equal(t, tsoutil.ComposeTSByTime(now.Add(-time.Second), 0), m.WALs[0].RecoveryTimeTick)

	local.EXPECT().GetMetricsIfLocal(mock.Anything).Unset()
	local.EXPECT().GetMetricsIfLocal(mock.Anything).Return(nil, errors.New("test"))
	m = getStreamingQuotaMetrics()
	assert.Nil(t, m)
}
