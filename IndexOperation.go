package vdatabase

const (
	OpEqual = iota * 30
	OpNotEqual
	OpGte
	OpGt
	OpLt
	OpLte
	OpIn
	OpIsNull
	OpIsNotNull
)

func NewIndexOperation(index *Index) *IndexOperation {
	return &IndexOperation{index}
}

type IndexOperation struct {
	index *Index
}

func (indexOperation *IndexOperation) CollectRowPointersByOperation(value interface{}, operation OperationType) []int {
	var rowPointers *[]int
	switch operation {
	case OpEqual:
		rowPointers = indexOperation.index.Find(value)
		break
	case OpNotEqual:
		rowPointers = indexOperation.index.FindNotEqual(value)
		break
	case OpIn:
		rowPointers = indexOperation.index.In(value)
		break
	case OpIsNull:
		rowPointers = indexOperation.index.IsNull()
		break
	case OpIsNotNull:
		rowPointers = indexOperation.index.IsNotNull()
		break
	case OpGte:
		rowPointers = indexOperation.index.GreaterEqual(value)
		break
	case OpGt:
		rowPointers = indexOperation.index.Greater(value)
		break
	case OpLte:
		rowPointers = indexOperation.index.LessEqual(value)
		break
	case OpLt:
		rowPointers = indexOperation.index.Less(value)
		break
	default:
		panic("Operation not found")
	}
	if nil == rowPointers {
		return []int{}
	}
	return *rowPointers
}
