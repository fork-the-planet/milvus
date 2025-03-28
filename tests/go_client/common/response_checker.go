package common

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/milvus-io/milvus/client/v2/column"
	"github.com/milvus-io/milvus/client/v2/entity"
	"github.com/milvus-io/milvus/client/v2/index"
	client "github.com/milvus-io/milvus/client/v2/milvusclient"
	"github.com/milvus-io/milvus/pkg/v2/log"
)

func CheckErr(t *testing.T, actualErr error, expErrNil bool, expErrorMsg ...string) {
	if expErrNil {
		require.NoError(t, actualErr)
	} else {
		require.Error(t, actualErr)
		switch len(expErrorMsg) {
		case 0:
			log.Fatal("expect error message should not be empty")
		case 1:
			require.ErrorContains(t, actualErr, expErrorMsg[0])
		default:
			contains := false
			for i := 0; i < len(expErrorMsg); i++ {
				if strings.Contains(actualErr.Error(), expErrorMsg[i]) {
					contains = true
				}
			}
			if !contains {
				t.Fatalf("CheckErr failed, actualErr doesn't contains any expErrorMsg, actual msg:%s", actualErr)
			}
		}
	}
}

// EqualColumn assert field data is equal of two columns
func EqualColumn(t *testing.T, columnA column.Column, columnB column.Column) {
	require.Equal(t, columnA.Name(), columnB.Name())
	require.Equal(t, columnA.Type(), columnB.Type())
	_type := columnA.Type()
	switch _type {
	case entity.FieldTypeBool:
		require.ElementsMatch(t, columnA.(*column.ColumnBool).Data(), columnB.(*column.ColumnBool).Data())
	case entity.FieldTypeInt8:
		require.ElementsMatch(t, columnA.(*column.ColumnInt8).Data(), columnB.(*column.ColumnInt8).Data())
	case entity.FieldTypeInt16:
		require.ElementsMatch(t, columnA.(*column.ColumnInt16).Data(), columnB.(*column.ColumnInt16).Data())
	case entity.FieldTypeInt32:
		require.ElementsMatch(t, columnA.(*column.ColumnInt32).Data(), columnB.(*column.ColumnInt32).Data())
	case entity.FieldTypeInt64:
		require.ElementsMatch(t, columnA.(*column.ColumnInt64).Data(), columnB.(*column.ColumnInt64).Data())
	case entity.FieldTypeFloat:
		require.ElementsMatch(t, columnA.(*column.ColumnFloat).Data(), columnB.(*column.ColumnFloat).Data())
	case entity.FieldTypeDouble:
		require.ElementsMatch(t, columnA.(*column.ColumnDouble).Data(), columnB.(*column.ColumnDouble).Data())
	case entity.FieldTypeVarChar:
		require.ElementsMatch(t, columnA.(*column.ColumnVarChar).Data(), columnB.(*column.ColumnVarChar).Data())
	case entity.FieldTypeJSON:
		log.Debug("data", zap.String("name", columnA.Name()), zap.Any("type", columnA.Type()), zap.Any("data", columnA.FieldData()))
		log.Debug("data", zap.String("name", columnB.Name()), zap.Any("type", columnB.Type()), zap.Any("data", columnB.FieldData()))
		require.Equal(t, reflect.TypeOf(columnA), reflect.TypeOf(columnB))
		switch _v := columnA.(type) {
		case *column.ColumnDynamic:
			require.ElementsMatch(t, columnA.(*column.ColumnDynamic).Data(), columnB.(*column.ColumnDynamic).Data())
		case *column.ColumnJSONBytes:
			require.ElementsMatch(t, columnA.(*column.ColumnJSONBytes).Data(), columnB.(*column.ColumnJSONBytes).Data())
		default:
			log.Warn("columnA type", zap.String("name", columnB.Name()), zap.Any("type", _v))
		}
	case entity.FieldTypeFloatVector:
		require.ElementsMatch(t, columnA.(*column.ColumnFloatVector).Data(), columnB.(*column.ColumnFloatVector).Data())
	case entity.FieldTypeBinaryVector:
		require.ElementsMatch(t, columnA.(*column.ColumnBinaryVector).Data(), columnB.(*column.ColumnBinaryVector).Data())
	case entity.FieldTypeFloat16Vector:
		require.ElementsMatch(t, columnA.(*column.ColumnFloat16Vector).Data(), columnB.(*column.ColumnFloat16Vector).Data())
	case entity.FieldTypeBFloat16Vector:
		require.ElementsMatch(t, columnA.(*column.ColumnBFloat16Vector).Data(), columnB.(*column.ColumnBFloat16Vector).Data())
	case entity.FieldTypeSparseVector:
		require.ElementsMatch(t, columnA.(*column.ColumnSparseFloatVector).Data(), columnB.(*column.ColumnSparseFloatVector).Data())
	case entity.FieldTypeArray:
		EqualArrayColumn(t, columnA, columnB)
	default:
		log.Info("Support column type is:", zap.Any("FieldType", []entity.FieldType{
			entity.FieldTypeBool,
			entity.FieldTypeInt8, entity.FieldTypeInt16, entity.FieldTypeInt32,
			entity.FieldTypeInt64, entity.FieldTypeFloat, entity.FieldTypeDouble, entity.FieldTypeString,
			entity.FieldTypeVarChar, entity.FieldTypeArray, entity.FieldTypeFloatVector, entity.FieldTypeBinaryVector,
		}))
	}
}

// EqualColumn assert field data is equal of two columns
func EqualArrayColumn(t *testing.T, columnA column.Column, columnB column.Column) {
	require.Equal(t, columnA.Name(), columnB.Name())
	require.IsType(t, columnA.Type(), entity.FieldTypeArray)
	require.IsType(t, columnB.Type(), entity.FieldTypeArray)
	switch _type := columnA.(type) {
	case *column.ColumnBoolArray:
		require.ElementsMatch(t, columnA.(*column.ColumnBoolArray).Data(), columnB.(*column.ColumnBoolArray).Data())
	case *column.ColumnInt8Array:
		require.ElementsMatch(t, columnA.(*column.ColumnInt8Array).Data(), columnB.(*column.ColumnInt8Array).Data())
	case *column.ColumnInt16Array:
		require.ElementsMatch(t, columnA.(*column.ColumnInt16Array).Data(), columnB.(*column.ColumnInt16Array).Data())
	case *column.ColumnInt32Array:
		require.ElementsMatch(t, columnA.(*column.ColumnInt32Array).Data(), columnB.(*column.ColumnInt32Array).Data())
	case *column.ColumnInt64Array:
		require.ElementsMatch(t, columnA.(*column.ColumnInt64Array).Data(), columnB.(*column.ColumnInt64Array).Data())
	case *column.ColumnFloatArray:
		require.ElementsMatch(t, columnA.(*column.ColumnFloatArray).Data(), columnB.(*column.ColumnFloatArray).Data())
	case *column.ColumnDoubleArray:
		require.ElementsMatch(t, columnA.(*column.ColumnDoubleArray).Data(), columnB.(*column.ColumnDoubleArray).Data())
	case *column.ColumnVarCharArray:
		require.ElementsMatch(t, columnA.(*column.ColumnVarCharArray).Data(), columnB.(*column.ColumnVarCharArray).Data())
	default:
		log.Debug("columnA type is", zap.Any("type", _type))
		log.Info("Support array element type is:", zap.Any("FieldType", []entity.FieldType{
			entity.FieldTypeBool, entity.FieldTypeInt8, entity.FieldTypeInt16,
			entity.FieldTypeInt32, entity.FieldTypeInt64, entity.FieldTypeFloat, entity.FieldTypeDouble, entity.FieldTypeVarChar,
		}))
	}
}

// CheckInsertResult check insert result, ids len (insert count), ids data (pks, but no auto ids)
func CheckInsertResult(t *testing.T, expIDs column.Column, insertRes client.InsertResult) {
	require.Equal(t, expIDs.Len(), insertRes.IDs.Len())
	require.Equal(t, expIDs.Len(), int(insertRes.InsertCount))
	actualIDs := insertRes.IDs
	switch expIDs.Type() {
	// pk field support int64 and varchar type
	case entity.FieldTypeInt64:
		require.ElementsMatch(t, actualIDs.(*column.ColumnInt64).Data(), expIDs.(*column.ColumnInt64).Data())
	case entity.FieldTypeVarChar:
		require.ElementsMatch(t, actualIDs.(*column.ColumnVarChar).Data(), expIDs.(*column.ColumnVarChar).Data())
	default:
		log.Info("The primary field only support ", zap.Any("type", []entity.FieldType{entity.FieldTypeInt64, entity.FieldTypeVarChar}))
	}
}

// CheckOutputFields check query output fields
func CheckOutputFields(t *testing.T, expFields []string, actualColumns []column.Column) {
	actualFields := make([]string, 0)
	for _, actualColumn := range actualColumns {
		actualFields = append(actualFields, actualColumn.Name())
	}
	log.Debug("CheckOutputFields", zap.Any("expFields", expFields), zap.Any("actualFields", actualFields))
	require.ElementsMatchf(t, expFields, actualFields, fmt.Sprintf("Expected search output fields: %v, actual: %v", expFields, actualFields))
}

// CheckSearchResult check search result, check nq, topk, ids, score
func CheckSearchResult(t *testing.T, actualSearchResults []client.ResultSet, expNq int, expTopK int) {
	require.Equalf(t, len(actualSearchResults), expNq, fmt.Sprintf("Expected nq=%d, actual SearchResultsLen=%d", expNq, len(actualSearchResults)))
	require.Len(t, actualSearchResults, expNq)
	for _, actualSearchResult := range actualSearchResults {
		require.Equalf(t, actualSearchResult.ResultCount, expTopK, fmt.Sprintf("Expected topK=%d, actual ResultCount=%d", expTopK, actualSearchResult.ResultCount))
		require.Equalf(t, actualSearchResult.IDs.Len(), expTopK, fmt.Sprintf("Expected topK=%d, actual IDsLen=%d", expTopK, actualSearchResult.IDs.Len()))
		require.Equalf(t, len(actualSearchResult.Scores), expTopK, fmt.Sprintf("Expected topK=%d, actual ScoresLen=%d", expTopK, len(actualSearchResult.Scores)))
	}
}

// CheckQueryResult check query result, column name, type and field
func CheckQueryResult(t *testing.T, expColumns []column.Column, actualColumns []column.Column) {
	require.Equal(t, len(actualColumns), len(expColumns),
		"The len of actual columns %d should greater or equal to the expected columns %d", len(actualColumns), len(expColumns))
	for _, expColumn := range expColumns {
		exist := false
		for _, actualColumn := range actualColumns {
			if expColumn.Name() == actualColumn.Name() {
				exist = true
				EqualColumn(t, expColumn, actualColumn)
			}
		}
		if !exist {
			log.Error("CheckQueryResult actualColumns no column", zap.String("name", expColumn.Name()))
		}
	}
}

// GenColumnDataOption -- create column data --
type checkIndexOpt struct {
	state            index.IndexState
	pendingIndexRows int64
	totalRows        int64
	indexedRows      int64
}

func TNewCheckIndexOpt(totalRows int64) *checkIndexOpt {
	return &checkIndexOpt{
		state:            IndexStateFinished,
		totalRows:        totalRows,
		pendingIndexRows: 0,
		indexedRows:      totalRows,
	}
}

func (opt *checkIndexOpt) TWithIndexState(state index.IndexState) *checkIndexOpt {
	opt.state = state
	return opt
}

func (opt *checkIndexOpt) TWithIndexRows(totalRows int64, indexedRows int64, pendingIndexRows int64) *checkIndexOpt {
	opt.totalRows = totalRows
	opt.indexedRows = indexedRows
	opt.pendingIndexRows = pendingIndexRows
	return opt
}

func CheckIndex(t *testing.T, actualIdxDesc client.IndexDescription, idx index.Index, opt *checkIndexOpt) {
	require.EqualValuesf(t, idx, actualIdxDesc.Index, "Actual index is not same with expected index")
	require.Equal(t, actualIdxDesc.TotalRows, actualIdxDesc.PendingIndexRows+actualIdxDesc.IndexedRows)
	if opt != nil {
		require.Equal(t, opt.totalRows, opt.pendingIndexRows+opt.indexedRows)
		require.Equal(t, opt.state, actualIdxDesc.State)
		require.Equal(t, opt.totalRows, actualIdxDesc.TotalRows)
		require.Equal(t, opt.indexedRows, actualIdxDesc.IndexedRows)
		require.Equal(t, opt.pendingIndexRows, actualIdxDesc.PendingIndexRows)
	}
}

func CheckTransfer(t *testing.T, actualRgs []*entity.ResourceGroupTransfer, expRgs []*entity.ResourceGroupTransfer) {
	if len(expRgs) == 0 {
		require.Len(t, actualRgs, 0)
	} else {
		_expRgs := make([]string, 0, len(expRgs))
		_actualRgs := make([]string, 0, len(actualRgs))
		for _, rg := range expRgs {
			_expRgs = append(_expRgs, rg.ResourceGroup)
		}
		for _, rg := range actualRgs {
			_actualRgs = append(_actualRgs, rg.ResourceGroup)
		}
		require.ElementsMatch(t, _expRgs, _actualRgs)
	}
}

func CheckResourceGroupConfig(t *testing.T, actualConfig *entity.ResourceGroupConfig, expConfig *entity.ResourceGroupConfig) {
	if expConfig.Requests.NodeNum != 0 {
		require.EqualValuesf(t, expConfig.Requests.NodeNum, actualConfig.Requests.NodeNum, "Requests.NodeNum mismatch")
	}

	if expConfig.Limits.NodeNum != 0 {
		require.EqualValuesf(t, expConfig.Limits.NodeNum, actualConfig.Limits.NodeNum, "Limits.NodeNum mismatch")
	}

	if expConfig.TransferFrom != nil {
		CheckTransfer(t, expConfig.TransferFrom, actualConfig.TransferFrom)
	}

	if expConfig.TransferTo != nil {
		CheckTransfer(t, expConfig.TransferTo, actualConfig.TransferTo)
	}
	if expConfig.NodeFilter.NodeLabels != nil {
		require.EqualValues(t, expConfig.NodeFilter, actualConfig.NodeFilter)
	}
}

func CheckResourceGroup(t *testing.T, actualRg *entity.ResourceGroup, expRg *entity.ResourceGroup) {
	require.EqualValues(t, expRg.Name, actualRg.Name, "ResourceGroup name mismatch")
	require.EqualValues(t, expRg.Capacity, actualRg.Capacity, "ResourceGroup capacity mismatch")
	if expRg.NumAvailableNode >= 0 {
		require.EqualValues(t, expRg.NumAvailableNode, len(actualRg.Nodes), "AvailableNodesNumber mismatch")
	}

	if expRg.Config != nil {
		CheckResourceGroupConfig(t, actualRg.Config, expRg.Config)
	}

	if expRg.Nodes != nil {
		require.ElementsMatch(t, expRg.Nodes, actualRg.Nodes, "Nodes count mismatch")
	}
}
