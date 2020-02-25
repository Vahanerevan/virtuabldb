package vdatabase

import (
	"fmt"
	"github.com/google/btree"
	"math"
	"reflect"
	"sync"
)

type vdbRow struct {
	dataItem interface{}
	values   map[string]interface{}
	ID       int
}

func (row vdbRow) Value(name string) interface{} {
	return row.values[name]
}

func (row vdbRow) Item() interface{} {
	return row.dataItem
}

type IndexType = int

type OperationType int

const MAX_SLOTS = math.MaxInt32

const (
	HashIndex IndexerType = iota
	BTreeIndex
)

const (
	TypeInt     DataType = "int"
	TypeString           = "string"
	TypeBool             = "bool"
	TypeFloat64          = "float64"
	TypeFloat32          = "float32"
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

func (virtualTable *VirtualTable) ListRows() []*vdbRow {
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

func (virtualTable *VirtualTable) CollectRowPointersByOperation(indexName string, value interface{}, operationType OperationType) []int {
	index := virtualTable.getIndex(indexName)
	return NewIndexOperation(index).CollectRowPointersByOperation(value, operationType)
}

func (virtualTable *VirtualTable) CollectRowPointersDataItem(rowPointers []int) (result []interface{}) {

	for _, rowId := range rowPointers {
		vp := virtualTable.GetRow(rowId)
		result = append(result, vp.dataItem)
	}
	return result
}

func (virtualTable *VirtualTable) CollectRowsByPointers(rowPointers []int) (result []vdbRow) {
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
	rowPointers := virtualTable.CollectRowPointersByOperation(indexName, value, operation)
	result := virtualTable.CollectRowPointersDataItem(rowPointers)
	return NewResultIterator(result)

	return nil
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
	item := reflect.ValueOf(inserted.dataItem)
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
	item := reflect.ValueOf(row.dataItem)
	if nil != row {
		for _, index := range virtualTable.indexMap {
			index.PurgeIndex(*row, item.FieldByName(index.Name).Interface())
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
