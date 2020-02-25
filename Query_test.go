package vdatabase

import (
	"reflect"
	"testing"
)

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