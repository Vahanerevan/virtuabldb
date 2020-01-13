package virtuabldb

import (
	"fmt"
	"github.com/google/btree"
	"math"
	"reflect"
	"sync"
)

type vdbRow struct {
	item   interface{}
	values map[string]interface{}
	ID     int
}

func (row vdbRow) Value(name string) interface{} {
	return row.values[name]
}

func (row vdbRow) Item() interface{} {
	return row.item
}

type IndexType = int

type OperationType int

const MAX_SLOTS = math.MaxInt32

const (
	HashIndex IndexerType = iota
	BTreeIndex
)

const (
	TYPE_INT     DataType = "int"
	TYPE_STRING           = "string"
	TYPE_BOOL             = "bool"
	TYPE_FLOAT64          = "float64"
	TYPE_FLOAT32          = "float32"
)

const (
	OP_EQUAL = iota * 30
	OP_NOT_EQUAL
	OP_GTE
	OP_GT
	OP_LT
	OP_LTE
	OP_IN
	OP_IS_NULL
	OP_IS_NOT_NULL
)

type VirtualTable struct {
	Name         string
	rows         []*vdbRow
	ForeignKey   []IndexOption
	IndexOptions IndexOptions
	indexMap     map[string]*Index
	primaryKey   *string
	sync.Mutex
}

func (virtualTable *VirtualTable) setPrimaryKey(primaryKey string) {
	virtualTable.primaryKey = &primaryKey
}

func (virtualTable *VirtualTable) ROWS() []*vdbRow {
	return virtualTable.rows
}

func (virtualTable *VirtualTable) Count() int {
	return len(virtualTable.rows)
}

func (virtualTable VirtualTable) getIndex(indexName string) *Index {
	value, ok := virtualTable.indexMap[indexName]
	if false == ok {
		panic(fmt.Sprintf("Index %s not found", indexName))
	}
	return value
}

func (virtualTable *VirtualTable) Where(indexName string, value interface{}) *ResultIterator {
	return virtualTable.findByOperation(indexName, value, OP_EQUAL)
}
func (virtualTable *VirtualTable) GreaterEqual(indexName string, value interface{}) *ResultIterator {
	return virtualTable.greaterEqual(indexName, value)
}
func (virtualTable *VirtualTable) LessThan(indexName string, value interface{}) *ResultIterator {
	return virtualTable.lessThan(indexName, value)
}

func (virtualTable *VirtualTable) getPointersByOperation(indexName string, value interface{}, operation OperationType) []int {
	var rowPointers *[]int
	index := virtualTable.getIndex(indexName)
	switch operation {
	case OP_EQUAL:
		rowPointers = index.Find(value)
		break
	case OP_NOT_EQUAL:
		rowPointers = index.FindNotEqual(value)
		break
	case OP_IN:
		rowPointers = index.In(value)
		break
	case OP_IS_NULL:
		rowPointers = index.IsNull()
		break
	case OP_IS_NOT_NULL:
		rowPointers = index.IsNotNull()
		break
	case OP_GTE:
		rowPointers = index.GreaterEqual(value)
		break
	case OP_GT:
		rowPointers = index.Greater(value)
		break
	case OP_LTE:
		rowPointers = index.LessEqual(value)
		break
	case OP_LT:
		rowPointers = index.Less(value)
		break
	default:
		panic("Operation not found")
	}
	if nil == rowPointers {
		return []int{}
	}
	return *rowPointers
}

func (virtualTable *VirtualTable) CollectRowPointersByOperation(indexName string, value interface{}, operationType OperationType) []int {
	return virtualTable.getPointersByOperation(indexName, value, operationType)
}

func (virtualTable *VirtualTable) CollectRowPointersInsertItems(rowPointers []int) (result []interface{}) {
	if nil != rowPointers {
		for _, rowId := range rowPointers {
			vp := virtualTable.GetRow(rowId)
			result = append(result, vp.item)
		}
	}
	return result
}

func (virtualTable *VirtualTable) CollectRowPointers(rowPointers []int) (result []vdbRow) {
	if nil != rowPointers {
		for _, rowId := range rowPointers {
			vp := virtualTable.GetRow(rowId)
			result = append(result, *vp)
		}
	}
	return result
}

func (virtualTable *VirtualTable) CollectNotNilPointers() (result []int) {
	for id, row := range virtualTable.rows {
		if nil != row {
			result = append(result, id)
		}
	}
	return result
}

func (virtualTable *VirtualTable) findByOperation(indexName string, value interface{}, operation OperationType) *ResultIterator {
	rowPointers := virtualTable.getPointersByOperation(indexName, value, operation)
	if nil != rowPointers {
		result := virtualTable.CollectRowPointersInsertItems(rowPointers)
		return NewResultIterator(result)

	}
	return nil
}

func (virtualTable VirtualTable) where(indexName string, value interface{}) *ResultIterator {
	return virtualTable.findByOperation(indexName, value, OP_EQUAL)
}

func (virtualTable *VirtualTable) greaterEqual(indexName string, value interface{}) *ResultIterator {
	return virtualTable.findByOperation(indexName, value, OP_GTE)
}

func (virtualTable *VirtualTable) lessThan(indexName string, value interface{}) *ResultIterator {
	return virtualTable.findByOperation(indexName, value, OP_LT)
}

func (virtualTable *VirtualTable) Initialize() {
	virtualTable.indexMap = make(map[string]*Index)
	virtualTable.rows = []*vdbRow{}
	for _, options := range virtualTable.IndexOptions {
		if true == options.IsPrimary {
			virtualTable.setPrimaryKey(options.Name)
		}
		virtualTable.indexMap[options.Name] = NewIndex(options)

	}

}

func (virtualTable *VirtualTable) insertInto(item interface{}) vdbRow {
	nextId := len(virtualTable.rows)
	r := vdbRow{item, make(map[string]interface{}), nextId}
	virtualTable.rows = append(virtualTable.rows, &r)
	return r
}

func (virtualTable *VirtualTable) GetRow(index int) *vdbRow {
	return virtualTable.rows[index]
}

func (virtualTable *VirtualTable) HasIndex(index string) bool {
	_, ok := virtualTable.indexMap[index]
	return ok
}

func (virtualTable *VirtualTable) createIndexes(inserted vdbRow) {
	item := reflect.ValueOf(inserted.item)
	for _, index := range virtualTable.indexMap {
		index.createIndex(inserted, item.FieldByName(index.Name).Interface())
	}

}

func (virtualTable *VirtualTable) DumpIndexes() {
	for _, index := range virtualTable.indexMap {
		index.Storage.Descend(func(i btree.Item) bool {
			return true
		})
	}
}

/**
CRUD
*/
func (virtualTable *VirtualTable) Insert(row interface{}) int {
	virtualTable.Lock()
	inserted := virtualTable.insertInto(row)
	virtualTable.createIndexes(inserted)
	virtualTable.Unlock()
	return inserted.ID
}

func (virtualTable *VirtualTable) Delete(rows ...int) {
	virtualTable.Lock()
	for _, index := range rows {
		virtualTable.deleteOne(index)
	}
	virtualTable.Unlock()
}

func (virtualTable *VirtualTable) deleteOne(rowId int) {
	row := virtualTable.GetRow(rowId)
	item := reflect.ValueOf(row.item)
	if nil != row {
		for _, index := range virtualTable.indexMap {
			index.removeIndex(*row, item.FieldByName(index.Name).Interface())
		}
	}
	virtualTable.rows[rowId] = nil
}

func (virtualTable *VirtualTable) Update(row interface{}) {
	//virtualTable.Rows = append(virtualTable.Rows, row)
}

/**
CRUD
*/
