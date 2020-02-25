package vdatabase

import (
	"fmt"
	"github.com/google/btree"
	"reflect"
	"sync"
)

type PointerStorage struct {
	Key      interface{}
	pointers *RowPointers
	DataType DataType
	mu       sync.Mutex
}

func (indexedItem PointerStorage) Pointers() RowPointers {
	return *indexedItem.pointers
}

func (indexedItem *PointerStorage) AddPointer(pointer RowPointer) {
	indexedItem.mu.Lock()
	if nil == indexedItem.pointers {
		indexedItem.pointers = NewRowPointers()
	}
	ptrs := indexedItem.pointers
	ptrs.Add(pointer)
	indexedItem.mu.Unlock()

}

func (indexedItem *PointerStorage) RemovePointer(pointer RowPointer) int {
	indexedItem.mu.Lock()
	if nil == indexedItem.pointers {
		return 0
	}
	indexedItem.pointers.Remove(pointer)
	indexedItem.mu.Unlock()

	return len(indexedItem.pointers.Pointers)
}

func (indexedItem PointerStorage) Less(than btree.Item) (result bool) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r, "TYPES", indexedItem.DataType ,"T = ", reflect.TypeOf(indexedItem.Key),indexedItem.Key, reflect.TypeOf(than.(*PointerStorage).Key),than.(*PointerStorage).Key)
		}
	}()
	return IndexResolvers.Less(indexedItem.DataType, indexedItem.Key, than.(*PointerStorage).Key)
}
