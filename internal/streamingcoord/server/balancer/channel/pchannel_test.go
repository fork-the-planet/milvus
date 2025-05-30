package channel

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/milvus-io/milvus/pkg/v2/proto/streamingpb"
	"github.com/milvus-io/milvus/pkg/v2/streaming/util/types"
)

func TestPChannel(t *testing.T) {
	ResetStaticPChannelStatsManager()
	RecoverPChannelStatsManager([]string{})

	pchannel := newPChannelMetaFromProto(&streamingpb.PChannelMeta{
		Channel: &streamingpb.PChannelInfo{
			Name: "test-channel",
			Term: 1,
		},
		Node: &streamingpb.StreamingNodeInfo{
			ServerId: 123,
		},
		State: streamingpb.PChannelMetaState_PCHANNEL_META_STATE_UNINITIALIZED,
	})
	assert.Equal(t, "test-channel", pchannel.Name())
	assert.Equal(t, int64(1), pchannel.CurrentTerm())
	assert.Equal(t, int64(123), pchannel.CurrentServerID())
	assert.Equal(t, streamingpb.PChannelMetaState_PCHANNEL_META_STATE_UNINITIALIZED, pchannel.State())
	assert.False(t, pchannel.IsAssigned())
	assert.Empty(t, pchannel.AssignHistories())
	assert.Equal(t, types.PChannelInfoAssigned{
		Channel: types.PChannelInfo{
			Name: "test-channel",
			Term: 1,
		},
		Node: types.StreamingNodeInfo{
			ServerID: 123,
		},
	}, pchannel.CurrentAssignment())

	pchannel = newPChannelMeta("test-channel", types.AccessModeRW)
	assert.Equal(t, "test-channel", pchannel.Name())
	assert.Equal(t, int64(1), pchannel.CurrentTerm())
	assert.Empty(t, pchannel.AssignHistories())
	assert.False(t, pchannel.IsAssigned())

	// Test CopyForWrite()
	mutablePChannel := pchannel.CopyForWrite()
	assert.NotNil(t, mutablePChannel)

	// Test AssignToServerID()
	newServerID := types.StreamingNodeInfo{
		ServerID: 456,
	}
	assert.True(t, mutablePChannel.TryAssignToServerID(types.AccessModeRW, newServerID))
	updatedChannelInfo := newPChannelMetaFromProto(mutablePChannel.IntoRawMeta())

	assert.Equal(t, "test-channel", pchannel.Name())
	assert.Equal(t, int64(1), pchannel.CurrentTerm())
	assert.Empty(t, pchannel.AssignHistories())

	assert.Equal(t, "test-channel", updatedChannelInfo.Name())
	assert.Equal(t, int64(2), updatedChannelInfo.CurrentTerm())
	assert.Equal(t, int64(456), updatedChannelInfo.CurrentServerID())
	assert.Empty(t, pchannel.AssignHistories())
	assert.False(t, updatedChannelInfo.IsAssigned())
	assert.Equal(t, streamingpb.PChannelMetaState_PCHANNEL_META_STATE_ASSIGNING, updatedChannelInfo.State())

	mutablePChannel = updatedChannelInfo.CopyForWrite()

	mutablePChannel.TryAssignToServerID(types.AccessModeRW, types.StreamingNodeInfo{ServerID: 789})
	updatedChannelInfo = newPChannelMetaFromProto(mutablePChannel.IntoRawMeta())
	assert.Equal(t, "test-channel", updatedChannelInfo.Name())
	assert.Equal(t, int64(3), updatedChannelInfo.CurrentTerm())
	assert.Equal(t, int64(789), updatedChannelInfo.CurrentServerID())
	assert.Len(t, updatedChannelInfo.AssignHistories(), 1)
	assert.Equal(t, "test-channel", updatedChannelInfo.AssignHistories()[0].Channel.Name)
	assert.Equal(t, int64(2), updatedChannelInfo.AssignHistories()[0].Channel.Term)
	assert.Equal(t, int64(456), updatedChannelInfo.AssignHistories()[0].Node.ServerID)
	assert.False(t, updatedChannelInfo.IsAssigned())
	assert.Equal(t, streamingpb.PChannelMetaState_PCHANNEL_META_STATE_ASSIGNING, updatedChannelInfo.State())

	// Test AssignToServerDone
	mutablePChannel = updatedChannelInfo.CopyForWrite()
	mutablePChannel.AssignToServerDone()
	updatedChannelInfo = newPChannelMetaFromProto(mutablePChannel.IntoRawMeta())
	assert.Equal(t, "test-channel", updatedChannelInfo.Name())
	assert.Equal(t, int64(3), updatedChannelInfo.CurrentTerm())
	assert.Equal(t, int64(789), updatedChannelInfo.CurrentServerID())
	assert.Len(t, updatedChannelInfo.AssignHistories(), 0)
	assert.True(t, updatedChannelInfo.IsAssigned())
	assert.Equal(t, streamingpb.PChannelMetaState_PCHANNEL_META_STATE_ASSIGNED, updatedChannelInfo.State())

	// Test reassigned
	mutablePChannel = updatedChannelInfo.CopyForWrite()
	assert.False(t, mutablePChannel.TryAssignToServerID(types.AccessModeRW, types.StreamingNodeInfo{ServerID: 789}))

	// Test MarkAsUnavailable
	mutablePChannel = updatedChannelInfo.CopyForWrite()
	mutablePChannel.MarkAsUnavailable(2)
	updatedChannelInfo = newPChannelMetaFromProto(mutablePChannel.IntoRawMeta())
	assert.True(t, updatedChannelInfo.IsAssigned())

	mutablePChannel = updatedChannelInfo.CopyForWrite()
	mutablePChannel.MarkAsUnavailable(3)
	updatedChannelInfo = newPChannelMetaFromProto(mutablePChannel.IntoRawMeta())
	assert.False(t, updatedChannelInfo.IsAssigned())
	assert.Equal(t, streamingpb.PChannelMetaState_PCHANNEL_META_STATE_UNAVAILABLE, updatedChannelInfo.State())

	// Test assign on unavailable
	mutablePChannel = updatedChannelInfo.CopyForWrite()
	assert.True(t, mutablePChannel.TryAssignToServerID(types.AccessModeRW, types.StreamingNodeInfo{ServerID: 789}))
	assert.Len(t, mutablePChannel.AssignHistories(), 1)
}
