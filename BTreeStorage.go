package vdatabase

import (
	"github.com/google/btree"
)

const BtreeDegree int = 9

func b2i(b bool) int8 {
	if b {
		return 1
	}
	return 0
}

func NewBTreeStorage() *BTreeStorage {
	return &BTreeStorage{BTree: btree.New(BtreeDegree)}
}

type BTreeStorage struct {
	*btree.BTree
}
