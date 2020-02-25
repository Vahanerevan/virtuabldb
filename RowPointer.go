package vdatabase

import (
	"sort"
)

func NewRowPointers() *RowPointers {
	return &RowPointers{
		Pointers: []int{},
	}
}

type RowPointers struct {
	Pointers []int
}

func (rowPointers *RowPointers) Add(pointer RowPointer) {
	data := append(rowPointers.Pointers, pointer)
	sort.Ints(data)
	rowPointers.Pointers = data
}

func (rowPointers *RowPointers) Remove(pointer RowPointer) {
	for index, value := range rowPointers.Pointers {
		if value == pointer {
			rowPointers.Pointers = append(rowPointers.Pointers[:index], rowPointers.Pointers[index+1:]...)
			break
		}
	}
}
