package vdatabase

import (
	"github.com/google/btree"
	"github.com/spf13/cast"
	"reflect"
)

//
//func b2i(b bool) int8 {
//	if b {
//		return 1
//	}
//	return 0
//}

func collect(result *[]int) func(item btree.Item) bool {
	return func(item btree.Item) bool {
		data := item.(*PointerStorage)
		ptrs := data.Pointers().Pointers
		*result = append(*result, ptrs...)
		return true
	}
}

func NewIndex(options IndexOption) *Index {
	return &Index{
		Storage:  NewBTreeStorage(),
		Name:     options.Name,
		DataType: options.Type,
		NullStorage: &PointerStorage{

		},
		Null: options.Null,
	}
}

type Index struct {
	Table       *VirtualTable
	Storage     *BTreeStorage
	Name        string
	DataType    DataType
	NullStorage *PointerStorage
	Null        bool
	//sync.Mutex
}

func (index *Index) AddToNullStore(pointerId int) {
	index.NullStorage.AddPointer(pointerId)
}

func (index *Index) Pointers(value interface{}) *RowPointers {
	item := PointerStorage{Key: value}
	btreeData := index.Storage.Get(item)
	if nil == btreeData {
		return nil
	}
	ptrs := btreeData.(*PointerStorage).Pointers()
	return &ptrs
}

func (index *Index) transformValue(value interface{}) *PointerStorage {

	if index.Null && reflect.ValueOf(value).IsNil() {
		return nil
	}
	value = IndexResolvers.Transform(index.DataType, value)
	return &PointerStorage{Key: value, DataType: index.DataType}
}

func (index *Index) All() *[]int {
	return index.IsNotNull()
}

func (index *Index) FindNotEqual(value interface{}) *[]int {

	gt := index.Greater(value)
	lt := index.Less(value)
	var data []int
	data = append(*lt, *gt...)
	return &data
}

func (index *Index) Find(value interface{}) *[]int {
	search := index.transformValue(value)
	btreeData := index.Storage.Get(search)
	if nil == btreeData {
		return nil
	}
	data := btreeData.(*PointerStorage)
	ptrs := data.Pointers().Pointers
	return &ptrs
}

func (index *Index) IsNotNull() *[]int {
	var pointers []int
	index.Storage.Ascend(func(i btree.Item) bool {
		data := i.(*PointerStorage)
		ptrs := data.Pointers().Pointers
		pointers = append(pointers, ptrs...)
		return true
	})
	return &pointers
}

func (index *Index) IsNull() *[]int {
	ptrs := index.NullStorage.Pointers().Pointers
	return &ptrs
}

func (index *Index) In(values interface{}) *[]int {

	list := cast.ToSlice(values)
	var pointers []int
	for _, value := range list {
		search := index.transformValue(value)
		btreeData := index.Storage.Get(search)

		if nil == btreeData {
			continue
		}
		data := btreeData.(*PointerStorage)
		ptrs := data.Pointers().Pointers
		pointers = append(pointers, ptrs...)

	}
	return &pointers
}

func (index *Index) GreaterEqual(value interface{}) *[]int {
	search := index.transformValue(value)
	var result []int
	index.Storage.AscendGreaterOrEqual(search, collect(&result))
	return &result
}

func (index *Index) Greater(value interface{}) *[]int {
	search := index.transformValue(value)
	var result []int
	index.Storage.DescendGreaterThan(search, collect(&result))
	return &result
}

func (index *Index) Less(value interface{}) *[]int {
	search := index.transformValue(value)
	var result []int
	index.Storage.AscendLessThan(search, collect(&result))
	return &result
}

func (index *Index) LessEqual(value interface{}) *[]int {
	search := index.transformValue(value)
	var result []int
	index.Storage.DescendLessOrEqual(search, collect(&result))
	return &result
}

func (index *Index) createIndex(inserted vdbRow, key interface{}) {
	item := index.transformValue(key)
	if nil == item && index.Null {
		index.NullStorage.AddPointer(inserted.ID)
		return
	}
	var btreeData = index.Storage.Get(item)
	if nil != btreeData {
		item = btreeData.(*PointerStorage)
	}
	item.AddPointer(inserted.ID)
	index.Storage.ReplaceOrInsert(item)

}

func (index *Index) PurgeIndex(inserted vdbRow, key interface{}) {
	item := index.transformValue(key)
	var btreeData = index.Storage.Get(item)
	if nil != btreeData {
		item = btreeData.(*PointerStorage)
	}

	cnt := item.RemovePointer(inserted.ID)
	if 0 == cnt {
		index.Storage.Delete(item)
	}
}
