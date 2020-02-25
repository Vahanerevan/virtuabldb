package vdatabase

import (
	"github.com/google/btree"
	"reflect"
	"sync"
	"testing"
)

func TestCreateTable(t *testing.T) {
	type args struct {
		table VirtualTable
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestDefineIndexResolver(t *testing.T) {
	type args struct {
		indexType DataType
		resolver  IndexResolver
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestIndexOperation_CollectRowPointersByOperation(t *testing.T) {
	type fields struct {
		index *Index
	}
	type args struct {
		value     interface{}
		operation OperationType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indexOperation := &IndexOperation{
				index: tt.fields.index,
			}
			if got := indexOperation.CollectRowPointersByOperation(tt.args.value, tt.args.operation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CollectRowPointersByOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_AddToNullStore(t *testing.T) {
	type fields struct {
		Table       *VirtualTable
		Storage     *BTreeStorage
		Name        string
		DataType    DataType
		NullStorage *PointerStorage
		Null        bool
	}
	type args struct {
		pointerId int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				Table:       tt.fields.Table,
				Storage:     tt.fields.Storage,
				Name:        tt.fields.Name,
				DataType:    tt.fields.DataType,
				NullStorage: tt.fields.NullStorage,
				Null:        tt.fields.Null,
			}
		})
	}
}

func TestIndex_All(t *testing.T) {
	type fields struct {
		Table       *VirtualTable
		Storage     *BTreeStorage
		Name        string
		DataType    DataType
		NullStorage *PointerStorage
		Null        bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *[]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				Table:       tt.fields.Table,
				Storage:     tt.fields.Storage,
				Name:        tt.fields.Name,
				DataType:    tt.fields.DataType,
				NullStorage: tt.fields.NullStorage,
				Null:        tt.fields.Null,
			}
			if got := index.All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_Find(t *testing.T) {
	type fields struct {
		Table       *VirtualTable
		Storage     *BTreeStorage
		Name        string
		DataType    DataType
		NullStorage *PointerStorage
		Null        bool
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				Table:       tt.fields.Table,
				Storage:     tt.fields.Storage,
				Name:        tt.fields.Name,
				DataType:    tt.fields.DataType,
				NullStorage: tt.fields.NullStorage,
				Null:        tt.fields.Null,
			}
			if got := index.Find(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_FindNotEqual(t *testing.T) {
	type fields struct {
		Table       *VirtualTable
		Storage     *BTreeStorage
		Name        string
		DataType    DataType
		NullStorage *PointerStorage
		Null        bool
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				Table:       tt.fields.Table,
				Storage:     tt.fields.Storage,
				Name:        tt.fields.Name,
				DataType:    tt.fields.DataType,
				NullStorage: tt.fields.NullStorage,
				Null:        tt.fields.Null,
			}
			if got := index.FindNotEqual(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindNotEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_Greater(t *testing.T) {
	type fields struct {
		Table       *VirtualTable
		Storage     *BTreeStorage
		Name        string
		DataType    DataType
		NullStorage *PointerStorage
		Null        bool
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				Table:       tt.fields.Table,
				Storage:     tt.fields.Storage,
				Name:        tt.fields.Name,
				DataType:    tt.fields.DataType,
				NullStorage: tt.fields.NullStorage,
				Null:        tt.fields.Null,
			}
			if got := index.Greater(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Greater() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_GreaterEqual(t *testing.T) {
	type fields struct {
		Table       *VirtualTable
		Storage     *BTreeStorage
		Name        string
		DataType    DataType
		NullStorage *PointerStorage
		Null        bool
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				Table:       tt.fields.Table,
				Storage:     tt.fields.Storage,
				Name:        tt.fields.Name,
				DataType:    tt.fields.DataType,
				NullStorage: tt.fields.NullStorage,
				Null:        tt.fields.Null,
			}
			if got := index.GreaterEqual(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GreaterEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_In(t *testing.T) {
	type fields struct {
		Table       *VirtualTable
		Storage     *BTreeStorage
		Name        string
		DataType    DataType
		NullStorage *PointerStorage
		Null        bool
	}
	type args struct {
		values interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				Table:       tt.fields.Table,
				Storage:     tt.fields.Storage,
				Name:        tt.fields.Name,
				DataType:    tt.fields.DataType,
				NullStorage: tt.fields.NullStorage,
				Null:        tt.fields.Null,
			}
			if got := index.In(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("In() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_IsNotNull(t *testing.T) {
	type fields struct {
		Table       *VirtualTable
		Storage     *BTreeStorage
		Name        string
		DataType    DataType
		NullStorage *PointerStorage
		Null        bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *[]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				Table:       tt.fields.Table,
				Storage:     tt.fields.Storage,
				Name:        tt.fields.Name,
				DataType:    tt.fields.DataType,
				NullStorage: tt.fields.NullStorage,
				Null:        tt.fields.Null,
			}
			if got := index.IsNotNull(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsNotNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_IsNull(t *testing.T) {
	type fields struct {
		Table       *VirtualTable
		Storage     *BTreeStorage
		Name        string
		DataType    DataType
		NullStorage *PointerStorage
		Null        bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *[]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				Table:       tt.fields.Table,
				Storage:     tt.fields.Storage,
				Name:        tt.fields.Name,
				DataType:    tt.fields.DataType,
				NullStorage: tt.fields.NullStorage,
				Null:        tt.fields.Null,
			}
			if got := index.IsNull(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_Less(t *testing.T) {
	type fields struct {
		Table       *VirtualTable
		Storage     *BTreeStorage
		Name        string
		DataType    DataType
		NullStorage *PointerStorage
		Null        bool
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				Table:       tt.fields.Table,
				Storage:     tt.fields.Storage,
				Name:        tt.fields.Name,
				DataType:    tt.fields.DataType,
				NullStorage: tt.fields.NullStorage,
				Null:        tt.fields.Null,
			}
			if got := index.Less(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_LessEqual(t *testing.T) {
	type fields struct {
		Table       *VirtualTable
		Storage     *BTreeStorage
		Name        string
		DataType    DataType
		NullStorage *PointerStorage
		Null        bool
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				Table:       tt.fields.Table,
				Storage:     tt.fields.Storage,
				Name:        tt.fields.Name,
				DataType:    tt.fields.DataType,
				NullStorage: tt.fields.NullStorage,
				Null:        tt.fields.Null,
			}
			if got := index.LessEqual(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LessEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_Pointers(t *testing.T) {
	type fields struct {
		Table       *VirtualTable
		Storage     *BTreeStorage
		Name        string
		DataType    DataType
		NullStorage *PointerStorage
		Null        bool
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RowPointers
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				Table:       tt.fields.Table,
				Storage:     tt.fields.Storage,
				Name:        tt.fields.Name,
				DataType:    tt.fields.DataType,
				NullStorage: tt.fields.NullStorage,
				Null:        tt.fields.Null,
			}
			if got := index.Pointers(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pointers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_PurgeIndex(t *testing.T) {
	type fields struct {
		Table       *VirtualTable
		Storage     *BTreeStorage
		Name        string
		DataType    DataType
		NullStorage *PointerStorage
		Null        bool
	}
	type args struct {
		inserted vdbRow
		key      interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				Table:       tt.fields.Table,
				Storage:     tt.fields.Storage,
				Name:        tt.fields.Name,
				DataType:    tt.fields.DataType,
				NullStorage: tt.fields.NullStorage,
				Null:        tt.fields.Null,
			}
		})
	}
}

func TestIndex_createIndex(t *testing.T) {
	type fields struct {
		Table       *VirtualTable
		Storage     *BTreeStorage
		Name        string
		DataType    DataType
		NullStorage *PointerStorage
		Null        bool
	}
	type args struct {
		inserted vdbRow
		key      interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				Table:       tt.fields.Table,
				Storage:     tt.fields.Storage,
				Name:        tt.fields.Name,
				DataType:    tt.fields.DataType,
				NullStorage: tt.fields.NullStorage,
				Null:        tt.fields.Null,
			}
		})
	}
}

func TestIndex_transformValue(t *testing.T) {
	type fields struct {
		Table       *VirtualTable
		Storage     *BTreeStorage
		Name        string
		DataType    DataType
		NullStorage *PointerStorage
		Null        bool
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PointerStorage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &Index{
				Table:       tt.fields.Table,
				Storage:     tt.fields.Storage,
				Name:        tt.fields.Name,
				DataType:    tt.fields.DataType,
				NullStorage: tt.fields.NullStorage,
				Null:        tt.fields.Null,
			}
			if got := index.transformValue(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transformValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBTreeStorage(t *testing.T) {
	tests := []struct {
		name string
		want *BTreeStorage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBTreeStorage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBTreeStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIndex(t *testing.T) {
	type args struct {
		options IndexOption
	}
	tests := []struct {
		name string
		args args
		want *Index
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIndex(tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIndexOperation(t *testing.T) {
	type args struct {
		index *Index
	}
	tests := []struct {
		name string
		args args
		want *IndexOperation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIndexOperation(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIndexOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewResultIterator(t *testing.T) {
	type args struct {
		items []interface{}
	}
	tests := []struct {
		name string
		args args
		want *ResultIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResultIterator(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResultIterator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRowPointers(t *testing.T) {
	tests := []struct {
		name string
		want *RowPointers
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRowPointers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRowPointers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPointerStorage_AddPointer(t *testing.T) {
	type fields struct {
		Key      interface{}
		pointers *RowPointers
		DataType DataType
		mu       sync.Mutex
	}
	type args struct {
		pointer RowPointer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indexedItem := &PointerStorage{
				Key:      tt.fields.Key,
				pointers: tt.fields.pointers,
				DataType: tt.fields.DataType,
				mu:       tt.fields.mu,
			}
		})
	}
}

func TestPointerStorage_Less(t *testing.T) {
	type fields struct {
		Key      interface{}
		pointers *RowPointers
		DataType DataType
		mu       sync.Mutex
	}
	type args struct {
		than btree.Item
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indexedItem := PointerStorage{
				Key:      tt.fields.Key,
				pointers: tt.fields.pointers,
				DataType: tt.fields.DataType,
				mu:       tt.fields.mu,
			}
			if gotResult := indexedItem.Less(tt.args.than); gotResult != tt.wantResult {
				t.Errorf("Less() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestPointerStorage_Pointers(t *testing.T) {
	type fields struct {
		Key      interface{}
		pointers *RowPointers
		DataType DataType
		mu       sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   RowPointers
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indexedItem := PointerStorage{
				Key:      tt.fields.Key,
				pointers: tt.fields.pointers,
				DataType: tt.fields.DataType,
				mu:       tt.fields.mu,
			}
			if got := indexedItem.Pointers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pointers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPointerStorage_RemovePointer(t *testing.T) {
	type fields struct {
		Key      interface{}
		pointers *RowPointers
		DataType DataType
		mu       sync.Mutex
	}
	type args struct {
		pointer RowPointer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indexedItem := &PointerStorage{
				Key:      tt.fields.Key,
				pointers: tt.fields.pointers,
				DataType: tt.fields.DataType,
				mu:       tt.fields.mu,
			}
			if got := indexedItem.RemovePointer(tt.args.pointer); got != tt.want {
				t.Errorf("RemovePointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintMemUsage(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestQuery(t *testing.T) {
	type args struct {
		tableName string
	}
	tests := []struct {
		name string
		args args
		want *QueryObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Query(tt.args.tableName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryManager_Equals(t *testing.T) {
	type fields struct {
		key   string
		query *QueryObject
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *QueryObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qm := &QueryManager{
				key:   tt.fields.key,
				query: tt.fields.query,
			}
			if got := qm.Equals(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryManager_Greater(t *testing.T) {
	type fields struct {
		key   string
		query *QueryObject
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *QueryObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qm := &QueryManager{
				key:   tt.fields.key,
				query: tt.fields.query,
			}
			if got := qm.Greater(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Greater() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryManager_GreaterEqual(t *testing.T) {
	type fields struct {
		key   string
		query *QueryObject
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *QueryObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qm := &QueryManager{
				key:   tt.fields.key,
				query: tt.fields.query,
			}
			if got := qm.GreaterEqual(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GreaterEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryManager_In(t *testing.T) {
	type fields struct {
		key   string
		query *QueryObject
	}
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *QueryObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qm := &QueryManager{
				key:   tt.fields.key,
				query: tt.fields.query,
			}
			if got := qm.In(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("In() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryManager_IsNotNull(t *testing.T) {
	type fields struct {
		key   string
		query *QueryObject
	}
	tests := []struct {
		name   string
		fields fields
		want   *QueryObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qm := &QueryManager{
				key:   tt.fields.key,
				query: tt.fields.query,
			}
			if got := qm.IsNotNull(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsNotNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryManager_IsNull(t *testing.T) {
	type fields struct {
		key   string
		query *QueryObject
	}
	tests := []struct {
		name   string
		fields fields
		want   *QueryObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qm := &QueryManager{
				key:   tt.fields.key,
				query: tt.fields.query,
			}
			if got := qm.IsNull(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryManager_Less(t *testing.T) {
	type fields struct {
		key   string
		query *QueryObject
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *QueryObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qm := &QueryManager{
				key:   tt.fields.key,
				query: tt.fields.query,
			}
			if got := qm.Less(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryManager_LessEqual(t *testing.T) {
	type fields struct {
		key   string
		query *QueryObject
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *QueryObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qm := &QueryManager{
				key:   tt.fields.key,
				query: tt.fields.query,
			}
			if got := qm.LessEqual(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LessEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryManager_NotEqual(t *testing.T) {
	type fields struct {
		key   string
		query *QueryObject
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *QueryObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qm := &QueryManager{
				key:   tt.fields.key,
				query: tt.fields.query,
			}
			if got := qm.NotEqual(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryObject_Delete(t *testing.T) {
	type fields struct {
		table      *VirtualTable
		conditions map[OperationType][]pair
		cache      map[int]reflect.Value
		order      []orderBy
		limit      *int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := &QueryObject{
				table:      tt.fields.table,
				conditions: tt.fields.conditions,
				cache:      tt.fields.cache,
				order:      tt.fields.order,
				limit:      tt.fields.limit,
			}
		})
	}
}

func TestQueryObject_Filter(t *testing.T) {
	type fields struct {
		table      *VirtualTable
		conditions map[OperationType][]pair
		cache      map[int]reflect.Value
		order      []orderBy
		limit      *int
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *QueryManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := &QueryObject{
				table:      tt.fields.table,
				conditions: tt.fields.conditions,
				cache:      tt.fields.cache,
				order:      tt.fields.order,
				limit:      tt.fields.limit,
			}
			if got := query.Filter(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryObject_First(t *testing.T) {
	type fields struct {
		table      *VirtualTable
		conditions map[OperationType][]pair
		cache      map[int]reflect.Value
		order      []orderBy
		limit      *int
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := &QueryObject{
				table:      tt.fields.table,
				conditions: tt.fields.conditions,
				cache:      tt.fields.cache,
				order:      tt.fields.order,
				limit:      tt.fields.limit,
			}
			if got := query.First(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("First() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryObject_Index(t *testing.T) {
	type fields struct {
		table      *VirtualTable
		conditions map[OperationType][]pair
		cache      map[int]reflect.Value
		order      []orderBy
		limit      *int
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *QueryManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := &QueryObject{
				table:      tt.fields.table,
				conditions: tt.fields.conditions,
				cache:      tt.fields.cache,
				order:      tt.fields.order,
				limit:      tt.fields.limit,
			}
			if got := query.Index(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryObject_Limit(t *testing.T) {
	type fields struct {
		table      *VirtualTable
		conditions map[OperationType][]pair
		cache      map[int]reflect.Value
		order      []orderBy
		limit      *int
	}
	type args struct {
		limit int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *QueryObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := &QueryObject{
				table:      tt.fields.table,
				conditions: tt.fields.conditions,
				cache:      tt.fields.cache,
				order:      tt.fields.order,
				limit:      tt.fields.limit,
			}
			if got := query.Limit(tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Limit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryObject_OrderBy(t *testing.T) {
	type fields struct {
		table      *VirtualTable
		conditions map[OperationType][]pair
		cache      map[int]reflect.Value
		order      []orderBy
		limit      *int
	}
	type args struct {
		name      string
		orderType OrderType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *QueryObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := &QueryObject{
				table:      tt.fields.table,
				conditions: tt.fields.conditions,
				cache:      tt.fields.cache,
				order:      tt.fields.order,
				limit:      tt.fields.limit,
			}
			if got := query.OrderBy(tt.args.name, tt.args.orderType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryObject_Result(t *testing.T) {
	type fields struct {
		table      *VirtualTable
		conditions map[OperationType][]pair
		cache      map[int]reflect.Value
		order      []orderBy
		limit      *int
	}
	tests := []struct {
		name   string
		fields fields
		want   *ResultIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := &QueryObject{
				table:      tt.fields.table,
				conditions: tt.fields.conditions,
				cache:      tt.fields.cache,
				order:      tt.fields.order,
				limit:      tt.fields.limit,
			}
			if got := query.Result(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Result() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryObject_RowCount(t *testing.T) {
	type fields struct {
		table      *VirtualTable
		conditions map[OperationType][]pair
		cache      map[int]reflect.Value
		order      []orderBy
		limit      *int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := &QueryObject{
				table:      tt.fields.table,
				conditions: tt.fields.conditions,
				cache:      tt.fields.cache,
				order:      tt.fields.order,
				limit:      tt.fields.limit,
			}
			if got := query.RowCount(); got != tt.want {
				t.Errorf("RowCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryObject_addCondition(t *testing.T) {
	type fields struct {
		table      *VirtualTable
		conditions map[OperationType][]pair
		cache      map[int]reflect.Value
		order      []orderBy
		limit      *int
	}
	type args struct {
		operationType OperationType
		key           string
		value         interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := &QueryObject{
				table:      tt.fields.table,
				conditions: tt.fields.conditions,
				cache:      tt.fields.cache,
				order:      tt.fields.order,
				limit:      tt.fields.limit,
			}
		})
	}
}

func TestQueryObject_cacheFoundItems(t *testing.T) {
	type fields struct {
		table      *VirtualTable
		conditions map[OperationType][]pair
		cache      map[int]reflect.Value
		order      []orderBy
		limit      *int
	}
	type args struct {
		rows []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := &QueryObject{
				table:      tt.fields.table,
				conditions: tt.fields.conditions,
				cache:      tt.fields.cache,
				order:      tt.fields.order,
				limit:      tt.fields.limit,
			}
		})
	}
}

func TestQueryObject_doOrderBy(t *testing.T) {
	type fields struct {
		table      *VirtualTable
		conditions map[OperationType][]pair
		cache      map[int]reflect.Value
		order      []orderBy
		limit      *int
	}
	type args struct {
		rows []vdbRow
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []vdbRow
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := &QueryObject{
				table:      tt.fields.table,
				conditions: tt.fields.conditions,
				cache:      tt.fields.cache,
				order:      tt.fields.order,
				limit:      tt.fields.limit,
			}
			if got := query.doOrderBy(tt.args.rows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doOrderBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryObject_filter(t *testing.T) {
	type fields struct {
		table      *VirtualTable
		conditions map[OperationType][]pair
		cache      map[int]reflect.Value
		order      []orderBy
		limit      *int
	}
	tests := []struct {
		name         string
		fields       fields
		wantRows     []interface{}
		wantPointers []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := &QueryObject{
				table:      tt.fields.table,
				conditions: tt.fields.conditions,
				cache:      tt.fields.cache,
				order:      tt.fields.order,
				limit:      tt.fields.limit,
			}
			gotRows, gotPointers := query.filter()
			if !reflect.DeepEqual(gotRows, tt.wantRows) {
				t.Errorf("filter() gotRows = %v, want %v", gotRows, tt.wantRows)
			}
			if !reflect.DeepEqual(gotPointers, tt.wantPointers) {
				t.Errorf("filter() gotPointers = %v, want %v", gotPointers, tt.wantPointers)
			}
		})
	}
}

func TestQueryObject_sortFunction(t *testing.T) {
	type fields struct {
		table      *VirtualTable
		conditions map[OperationType][]pair
		cache      map[int]reflect.Value
		order      []orderBy
		limit      *int
	}
	type args struct {
		orderType OrderType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   func(x, y interface{}, dataType DataType) bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := QueryObject{
				table:      tt.fields.table,
				conditions: tt.fields.conditions,
				cache:      tt.fields.cache,
				order:      tt.fields.order,
				limit:      tt.fields.limit,
			}
			if got := query.sortFunction(tt.args.orderType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResultIterator_Count(t *testing.T) {
	type fields struct {
		cursor int
		items  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultIterator := ResultIterator{
				cursor: tt.fields.cursor,
				items:  tt.fields.items,
			}
			if got := resultIterator.Count(); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResultIterator_First(t *testing.T) {
	type fields struct {
		cursor int
		items  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultIterator := &ResultIterator{
				cursor: tt.fields.cursor,
				items:  tt.fields.items,
			}
			if got := resultIterator.First(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("First() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResultIterator_Item(t *testing.T) {
	type fields struct {
		cursor int
		items  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultIterator := &ResultIterator{
				cursor: tt.fields.cursor,
				items:  tt.fields.items,
			}
			if got := resultIterator.Item(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Item() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResultIterator_Iterate(t *testing.T) {
	type fields struct {
		cursor int
		items  []interface{}
	}
	type args struct {
		callable func(interface{})
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultIterator := &ResultIterator{
				cursor: tt.fields.cursor,
				items:  tt.fields.items,
			}
		})
	}
}

func TestResultIterator_Last(t *testing.T) {
	type fields struct {
		cursor int
		items  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultIterator := &ResultIterator{
				cursor: tt.fields.cursor,
				items:  tt.fields.items,
			}
			if got := resultIterator.Last(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Last() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResultIterator_Length(t *testing.T) {
	type fields struct {
		cursor int
		items  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultIterator := ResultIterator{
				cursor: tt.fields.cursor,
				items:  tt.fields.items,
			}
			if got := resultIterator.Length(); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResultIterator_Next(t *testing.T) {
	type fields struct {
		cursor int
		items  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultIterator := &ResultIterator{
				cursor: tt.fields.cursor,
				items:  tt.fields.items,
			}
			if got := resultIterator.Next(); got != tt.want {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResultIterator_Reset(t *testing.T) {
	type fields struct {
		cursor int
		items  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultIterator := &ResultIterator{
				cursor: tt.fields.cursor,
				items:  tt.fields.items,
			}
		})
	}
}

func TestResultIterator_SetItems(t *testing.T) {
	type fields struct {
		cursor int
		items  []interface{}
	}
	type args struct {
		_items []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ResultIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultIterator := &ResultIterator{
				cursor: tt.fields.cursor,
				items:  tt.fields.items,
			}
			if got := resultIterator.SetItems(tt.args._items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResultIterator_do(t *testing.T) {
	type fields struct {
		cursor int
		items  []interface{}
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultIterator := &ResultIterator{
				cursor: tt.fields.cursor,
				items:  tt.fields.items,
			}
			if got := resultIterator.do(tt.args.value); got != tt.want {
				t.Errorf("do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResultIterator_get(t *testing.T) {
	type fields struct {
		cursor int
		items  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultIterator := ResultIterator{
				cursor: tt.fields.cursor,
				items:  tt.fields.items,
			}
			if got := resultIterator.get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResultIterator_has(t *testing.T) {
	type fields struct {
		cursor int
		items  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultIterator := ResultIterator{
				cursor: tt.fields.cursor,
				items:  tt.fields.items,
			}
			if got := resultIterator.has(); got != tt.want {
				t.Errorf("has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResultIterator_hasIndex(t *testing.T) {
	type fields struct {
		cursor int
		items  []interface{}
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultIterator := ResultIterator{
				cursor: tt.fields.cursor,
				items:  tt.fields.items,
			}
			if got := resultIterator.hasIndex(tt.args.index); got != tt.want {
				t.Errorf("hasIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRowPointers_Add(t *testing.T) {
	type fields struct {
		Pointers []int
	}
	type args struct {
		pointer RowPointer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rowPointers := &RowPointers{
				Pointers: tt.fields.Pointers,
			}
		})
	}
}

func TestRowPointers_Remove(t *testing.T) {
	type fields struct {
		Pointers []int
	}
	type args struct {
		pointer RowPointer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rowPointers := &RowPointers{
				Pointers: tt.fields.Pointers,
			}
		})
	}
}

func TestSchema_Table(t *testing.T) {
	type fields struct {
		tables map[string]*VirtualTable
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *VirtualTable
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schema := &Schema{
				tables: tt.fields.tables,
			}
			if got := schema.Table(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_addTable(t *testing.T) {
	type fields struct {
		tables map[string]*VirtualTable
	}
	type args struct {
		table *VirtualTable
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schema := &Schema{
				tables: tt.fields.tables,
			}
		})
	}
}

func TestSchema_getTable(t *testing.T) {
	type fields struct {
		tables map[string]*VirtualTable
	}
	type args struct {
		tableName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *VirtualTable
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schema := &Schema{
				tables: tt.fields.tables,
			}
			if got := schema.getTable(tt.args.tableName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *VirtualTable
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Table(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVirtualTable_CollectNotNilPointers(t *testing.T) {
	type fields struct {
		Name         string
		rows         []*vdbRow
		ForeignKey   []IndexOption
		IndexOptions IndexOptions
		indexMap     map[string]*Index
		primaryKey   *string
		Mutex        sync.Mutex
	}
	tests := []struct {
		name       string
		fields     fields
		wantResult []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			virtualTable := &VirtualTable{
				Name:         tt.fields.Name,
				rows:         tt.fields.rows,
				ForeignKey:   tt.fields.ForeignKey,
				IndexOptions: tt.fields.IndexOptions,
				indexMap:     tt.fields.indexMap,
				primaryKey:   tt.fields.primaryKey,
				Mutex:        tt.fields.Mutex,
			}
			if gotResult := virtualTable.CollectNotNilPointers(); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("CollectNotNilPointers() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestVirtualTable_CollectRowPointersByOperation(t *testing.T) {
	type fields struct {
		Name         string
		rows         []*vdbRow
		ForeignKey   []IndexOption
		IndexOptions IndexOptions
		indexMap     map[string]*Index
		primaryKey   *string
		Mutex        sync.Mutex
	}
	type args struct {
		indexName     string
		value         interface{}
		operationType OperationType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			virtualTable := &VirtualTable{
				Name:         tt.fields.Name,
				rows:         tt.fields.rows,
				ForeignKey:   tt.fields.ForeignKey,
				IndexOptions: tt.fields.IndexOptions,
				indexMap:     tt.fields.indexMap,
				primaryKey:   tt.fields.primaryKey,
				Mutex:        tt.fields.Mutex,
			}
			if got := virtualTable.CollectRowPointersByOperation(tt.args.indexName, tt.args.value, tt.args.operationType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CollectRowPointersByOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVirtualTable_CollectRowPointersDataItem(t *testing.T) {
	type fields struct {
		Name         string
		rows         []*vdbRow
		ForeignKey   []IndexOption
		IndexOptions IndexOptions
		indexMap     map[string]*Index
		primaryKey   *string
		Mutex        sync.Mutex
	}
	type args struct {
		rowPointers []int
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			virtualTable := &VirtualTable{
				Name:         tt.fields.Name,
				rows:         tt.fields.rows,
				ForeignKey:   tt.fields.ForeignKey,
				IndexOptions: tt.fields.IndexOptions,
				indexMap:     tt.fields.indexMap,
				primaryKey:   tt.fields.primaryKey,
				Mutex:        tt.fields.Mutex,
			}
			if gotResult := virtualTable.CollectRowPointersDataItem(tt.args.rowPointers); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("CollectRowPointersDataItem() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestVirtualTable_CollectRowsByPointers(t *testing.T) {
	type fields struct {
		Name         string
		rows         []*vdbRow
		ForeignKey   []IndexOption
		IndexOptions IndexOptions
		indexMap     map[string]*Index
		primaryKey   *string
		Mutex        sync.Mutex
	}
	type args struct {
		rowPointers []int
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult []vdbRow
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			virtualTable := &VirtualTable{
				Name:         tt.fields.Name,
				rows:         tt.fields.rows,
				ForeignKey:   tt.fields.ForeignKey,
				IndexOptions: tt.fields.IndexOptions,
				indexMap:     tt.fields.indexMap,
				primaryKey:   tt.fields.primaryKey,
				Mutex:        tt.fields.Mutex,
			}
			if gotResult := virtualTable.CollectRowsByPointers(tt.args.rowPointers); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("CollectRowsByPointers() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestVirtualTable_Count(t *testing.T) {
	type fields struct {
		Name         string
		rows         []*vdbRow
		ForeignKey   []IndexOption
		IndexOptions IndexOptions
		indexMap     map[string]*Index
		primaryKey   *string
		Mutex        sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			virtualTable := &VirtualTable{
				Name:         tt.fields.Name,
				rows:         tt.fields.rows,
				ForeignKey:   tt.fields.ForeignKey,
				IndexOptions: tt.fields.IndexOptions,
				indexMap:     tt.fields.indexMap,
				primaryKey:   tt.fields.primaryKey,
				Mutex:        tt.fields.Mutex,
			}
			if got := virtualTable.Count(); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVirtualTable_Delete(t *testing.T) {
	type fields struct {
		Name         string
		rows         []*vdbRow
		ForeignKey   []IndexOption
		IndexOptions IndexOptions
		indexMap     map[string]*Index
		primaryKey   *string
		Mutex        sync.Mutex
	}
	type args struct {
		rows []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			virtualTable := &VirtualTable{
				Name:         tt.fields.Name,
				rows:         tt.fields.rows,
				ForeignKey:   tt.fields.ForeignKey,
				IndexOptions: tt.fields.IndexOptions,
				indexMap:     tt.fields.indexMap,
				primaryKey:   tt.fields.primaryKey,
				Mutex:        tt.fields.Mutex,
			}
		})
	}
}

func TestVirtualTable_DumpIndexes(t *testing.T) {
	type fields struct {
		Name         string
		rows         []*vdbRow
		ForeignKey   []IndexOption
		IndexOptions IndexOptions
		indexMap     map[string]*Index
		primaryKey   *string
		Mutex        sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			virtualTable := &VirtualTable{
				Name:         tt.fields.Name,
				rows:         tt.fields.rows,
				ForeignKey:   tt.fields.ForeignKey,
				IndexOptions: tt.fields.IndexOptions,
				indexMap:     tt.fields.indexMap,
				primaryKey:   tt.fields.primaryKey,
				Mutex:        tt.fields.Mutex,
			}
		})
	}
}

func TestVirtualTable_GetRow(t *testing.T) {
	type fields struct {
		Name         string
		rows         []*vdbRow
		ForeignKey   []IndexOption
		IndexOptions IndexOptions
		indexMap     map[string]*Index
		primaryKey   *string
		Mutex        sync.Mutex
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *vdbRow
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			virtualTable := &VirtualTable{
				Name:         tt.fields.Name,
				rows:         tt.fields.rows,
				ForeignKey:   tt.fields.ForeignKey,
				IndexOptions: tt.fields.IndexOptions,
				indexMap:     tt.fields.indexMap,
				primaryKey:   tt.fields.primaryKey,
				Mutex:        tt.fields.Mutex,
			}
			if got := virtualTable.GetRow(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVirtualTable_HasIndex(t *testing.T) {
	type fields struct {
		Name         string
		rows         []*vdbRow
		ForeignKey   []IndexOption
		IndexOptions IndexOptions
		indexMap     map[string]*Index
		primaryKey   *string
		Mutex        sync.Mutex
	}
	type args struct {
		index string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			virtualTable := &VirtualTable{
				Name:         tt.fields.Name,
				rows:         tt.fields.rows,
				ForeignKey:   tt.fields.ForeignKey,
				IndexOptions: tt.fields.IndexOptions,
				indexMap:     tt.fields.indexMap,
				primaryKey:   tt.fields.primaryKey,
				Mutex:        tt.fields.Mutex,
			}
			if got := virtualTable.HasIndex(tt.args.index); got != tt.want {
				t.Errorf("HasIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVirtualTable_Initialize(t *testing.T) {
	type fields struct {
		Name         string
		rows         []*vdbRow
		ForeignKey   []IndexOption
		IndexOptions IndexOptions
		indexMap     map[string]*Index
		primaryKey   *string
		Mutex        sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			virtualTable := &VirtualTable{
				Name:         tt.fields.Name,
				rows:         tt.fields.rows,
				ForeignKey:   tt.fields.ForeignKey,
				IndexOptions: tt.fields.IndexOptions,
				indexMap:     tt.fields.indexMap,
				primaryKey:   tt.fields.primaryKey,
				Mutex:        tt.fields.Mutex,
			}
		})
	}
}

func TestVirtualTable_Insert(t *testing.T) {
	type fields struct {
		Name         string
		rows         []*vdbRow
		ForeignKey   []IndexOption
		IndexOptions IndexOptions
		indexMap     map[string]*Index
		primaryKey   *string
		Mutex        sync.Mutex
	}
	type args struct {
		row interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			virtualTable := &VirtualTable{
				Name:         tt.fields.Name,
				rows:         tt.fields.rows,
				ForeignKey:   tt.fields.ForeignKey,
				IndexOptions: tt.fields.IndexOptions,
				indexMap:     tt.fields.indexMap,
				primaryKey:   tt.fields.primaryKey,
				Mutex:        tt.fields.Mutex,
			}
			if got := virtualTable.Insert(tt.args.row); got != tt.want {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVirtualTable_ListRows(t *testing.T) {
	type fields struct {
		Name         string
		rows         []*vdbRow
		ForeignKey   []IndexOption
		IndexOptions IndexOptions
		indexMap     map[string]*Index
		primaryKey   *string
		Mutex        sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   []*vdbRow
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			virtualTable := &VirtualTable{
				Name:         tt.fields.Name,
				rows:         tt.fields.rows,
				ForeignKey:   tt.fields.ForeignKey,
				IndexOptions: tt.fields.IndexOptions,
				indexMap:     tt.fields.indexMap,
				primaryKey:   tt.fields.primaryKey,
				Mutex:        tt.fields.Mutex,
			}
			if got := virtualTable.ListRows(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListRows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVirtualTable_Update(t *testing.T) {
	type fields struct {
		Name         string
		rows         []*vdbRow
		ForeignKey   []IndexOption
		IndexOptions IndexOptions
		indexMap     map[string]*Index
		primaryKey   *string
		Mutex        sync.Mutex
	}
	type args struct {
		row interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			virtualTable := &VirtualTable{
				Name:         tt.fields.Name,
				rows:         tt.fields.rows,
				ForeignKey:   tt.fields.ForeignKey,
				IndexOptions: tt.fields.IndexOptions,
				indexMap:     tt.fields.indexMap,
				primaryKey:   tt.fields.primaryKey,
				Mutex:        tt.fields.Mutex,
			}
		})
	}
}

func TestVirtualTable_createIndexes(t *testing.T) {
	type fields struct {
		Name         string
		rows         []*vdbRow
		ForeignKey   []IndexOption
		IndexOptions IndexOptions
		indexMap     map[string]*Index
		primaryKey   *string
		Mutex        sync.Mutex
	}
	type args struct {
		inserted vdbRow
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			virtualTable := &VirtualTable{
				Name:         tt.fields.Name,
				rows:         tt.fields.rows,
				ForeignKey:   tt.fields.ForeignKey,
				IndexOptions: tt.fields.IndexOptions,
				indexMap:     tt.fields.indexMap,
				primaryKey:   tt.fields.primaryKey,
				Mutex:        tt.fields.Mutex,
			}
		})
	}
}

func TestVirtualTable_deleteOne(t *testing.T) {
	type fields struct {
		Name         string
		rows         []*vdbRow
		ForeignKey   []IndexOption
		IndexOptions IndexOptions
		indexMap     map[string]*Index
		primaryKey   *string
		Mutex        sync.Mutex
	}
	type args struct {
		rowId int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			virtualTable := &VirtualTable{
				Name:         tt.fields.Name,
				rows:         tt.fields.rows,
				ForeignKey:   tt.fields.ForeignKey,
				IndexOptions: tt.fields.IndexOptions,
				indexMap:     tt.fields.indexMap,
				primaryKey:   tt.fields.primaryKey,
				Mutex:        tt.fields.Mutex,
			}
		})
	}
}

func TestVirtualTable_findByOperation(t *testing.T) {
	type fields struct {
		Name         string
		rows         []*vdbRow
		ForeignKey   []IndexOption
		IndexOptions IndexOptions
		indexMap     map[string]*Index
		primaryKey   *string
		Mutex        sync.Mutex
	}
	type args struct {
		indexName string
		value     interface{}
		operation OperationType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ResultIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			virtualTable := &VirtualTable{
				Name:         tt.fields.Name,
				rows:         tt.fields.rows,
				ForeignKey:   tt.fields.ForeignKey,
				IndexOptions: tt.fields.IndexOptions,
				indexMap:     tt.fields.indexMap,
				primaryKey:   tt.fields.primaryKey,
				Mutex:        tt.fields.Mutex,
			}
			if got := virtualTable.findByOperation(tt.args.indexName, tt.args.value, tt.args.operation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findByOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVirtualTable_getIndex(t *testing.T) {
	type fields struct {
		Name         string
		rows         []*vdbRow
		ForeignKey   []IndexOption
		IndexOptions IndexOptions
		indexMap     map[string]*Index
		primaryKey   *string
		Mutex        sync.Mutex
	}
	type args struct {
		indexName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Index
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			virtualTable := VirtualTable{
				Name:         tt.fields.Name,
				rows:         tt.fields.rows,
				ForeignKey:   tt.fields.ForeignKey,
				IndexOptions: tt.fields.IndexOptions,
				indexMap:     tt.fields.indexMap,
				primaryKey:   tt.fields.primaryKey,
				Mutex:        tt.fields.Mutex,
			}
			if got := virtualTable.getIndex(tt.args.indexName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVirtualTable_insertInto(t *testing.T) {
	type fields struct {
		Name         string
		rows         []*vdbRow
		ForeignKey   []IndexOption
		IndexOptions IndexOptions
		indexMap     map[string]*Index
		primaryKey   *string
		Mutex        sync.Mutex
	}
	type args struct {
		item interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   vdbRow
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			virtualTable := &VirtualTable{
				Name:         tt.fields.Name,
				rows:         tt.fields.rows,
				ForeignKey:   tt.fields.ForeignKey,
				IndexOptions: tt.fields.IndexOptions,
				indexMap:     tt.fields.indexMap,
				primaryKey:   tt.fields.primaryKey,
				Mutex:        tt.fields.Mutex,
			}
			if got := virtualTable.insertInto(tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertInto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVirtualTable_setPrimaryKey(t *testing.T) {
	type fields struct {
		Name         string
		rows         []*vdbRow
		ForeignKey   []IndexOption
		IndexOptions IndexOptions
		indexMap     map[string]*Index
		primaryKey   *string
		Mutex        sync.Mutex
	}
	type args struct {
		primaryKey string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			virtualTable := &VirtualTable{
				Name:         tt.fields.Name,
				rows:         tt.fields.rows,
				ForeignKey:   tt.fields.ForeignKey,
				IndexOptions: tt.fields.IndexOptions,
				indexMap:     tt.fields.indexMap,
				primaryKey:   tt.fields.primaryKey,
				Mutex:        tt.fields.Mutex,
			}
		})
	}
}

func Test_b2i(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b2i(tt.args.b); got != tt.want {
				t.Errorf("b2i() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bToMb(t *testing.T) {
	type args struct {
		b uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bToMb(tt.args.b); got != tt.want {
				t.Errorf("bToMb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_collect(t *testing.T) {
	type args struct {
		result *[]int
	}
	tests := []struct {
		name string
		args args
		want func(item btree.Item) bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := collect(tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("collect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_collectRowsData(t *testing.T) {
	type args struct {
		rows []vdbRow
	}
	tests := []struct {
		name    string
		args    args
		wantRes []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := collectRowsData(tt.args.rows); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("collectRowsData() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_dataTypeResolver_Less(t *testing.T) {
	type fields struct {
		Resolvers map[DataType]IndexResolver
	}
	type args struct {
		dataType DataType
		item     interface{}
		than     interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dte := &dataTypeResolver{
				Resolvers: tt.fields.Resolvers,
			}
			if got := dte.Less(tt.args.dataType, tt.args.item, tt.args.than); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataTypeResolver_OnNull(t *testing.T) {
	type fields struct {
		Resolvers map[DataType]IndexResolver
	}
	type args struct {
		dataType DataType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dte := &dataTypeResolver{
				Resolvers: tt.fields.Resolvers,
			}
			if got := dte.OnNull(tt.args.dataType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OnNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataTypeResolver_Transform(t *testing.T) {
	type fields struct {
		Resolvers map[DataType]IndexResolver
	}
	type args struct {
		dataType DataType
		item     interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dte := &dataTypeResolver{
				Resolvers: tt.fields.Resolvers,
			}
			if got := dte.Transform(tt.args.dataType, tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resetQuery(t *testing.T) {
	type args struct {
		q *QueryObject
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_vdbRow_Item(t *testing.T) {
	type fields struct {
		dataItem interface{}
		values   map[string]interface{}
		ID       int
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			row := vdbRow{
				dataItem: tt.fields.dataItem,
				values:   tt.fields.values,
				ID:       tt.fields.ID,
			}
			if got := row.Item(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Item() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vdbRow_Value(t *testing.T) {
	type fields struct {
		dataItem interface{}
		values   map[string]interface{}
		ID       int
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			row := vdbRow{
				dataItem: tt.fields.dataItem,
				values:   tt.fields.values,
				ID:       tt.fields.ID,
			}
			if got := row.Value(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}