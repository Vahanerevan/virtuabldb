package virtuabldb

import (
	"github.com/google/btree"
)

func b2i(b bool) int8 {
	if b {
		return 1
	}
	return 0
}


func NewBTreeStorage() *BTreeStorage {
	return &BTreeStorage{BTree: btree.New(9)}
}

type BTreeStorage struct {
	*btree.BTree
}

