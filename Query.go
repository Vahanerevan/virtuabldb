package vdatabase

import (
	"github.com/juliangruber/go-intersect"
	"github.com/spf13/cast"
	"reflect"
	"sort"
)

func Query(tableName string) *QueryObject {
	return &QueryObject{
		table:      Table(tableName),
		conditions: map[OperationType][]pair{},
	}
}

func collectRowsData(rows []vdbRow) (res []interface{}) {
	for _, v := range rows {
		res = append(res, v.Item())
	}
	return res
}

type orderBy struct {
	Key  string
	Type OrderType
}

type QueryObject struct {
	table      *VirtualTable
	conditions map[OperationType][]pair
	cache      map[int]reflect.Value
	order      []orderBy
	limit      *int
}

func (query *QueryObject) reset() {
	query.conditions = map[OperationType][]pair{}
}

func (query *QueryObject) RowCount() int {
	return query.table.Count()
}

func (query *QueryObject) Index(name string) *QueryManager {
	return &QueryManager{
		key:   name,
		query: query,
	}
}

func (query *QueryObject) Filter(name string) *QueryManager {
	return query.Index(name)
}

func (query *QueryObject) addCondition(operationType OperationType, key string, value interface{}) {
	query.conditions[operationType] = append(query.conditions[operationType], pair{
		Key:   key,
		Value: value,
	})
}

func (query *QueryObject) filter() (rows []interface{}, pointers []int) {

	if 0 == len(query.conditions) {
		pointers = query.table.CollectNotNilPointers()
	} else {
		var isFirstIteration = true
		isDone := false
		for operationType, pairs := range query.conditions {
			if isDone {
				break
			}
			for _, p := range pairs {
				result := query.table.CollectRowPointersByOperation(p.Key, p.Value, operationType)
				if nil != result {
					if true == isFirstIteration {
						pointers = result
						isFirstIteration = false
					} else {
						res := intersect.Simple(pointers, result)
						intersections := cast.ToIntSlice(res)
						pointers = intersections
					}
				} else {
					pointers = []int{}
					isDone = true
					break
				}
			}
		}
	}

	if nil != query.limit {
		if len(pointers) >= *query.limit {
			pointers = pointers[0:*query.limit]
		}
	}

	vdbRows := query.table.CollectRowsByPointers(pointers)

	//query.cacheFoundItems(vdbRows)

	if 0 != len(query.order) {
		vdbRows = query.doOrderBy(vdbRows)
	}

	rows = collectRowsData(vdbRows)
	return rows, pointers
}

func (query *QueryObject) Result() *ResultIterator {
	rows, _ := query.filter()
	query.reset()
	return NewResultIterator(rows)
}
func (query *QueryObject) First() interface{} {
	rows, _ := query.filter()
	query.reset()

	if len(rows) > 0 {
		return rows[0]
	}
	return nil
}

func (query *QueryObject) OrderBy(name string, orderType OrderType) *QueryObject {
	order := orderBy{
		Key:  name,
		Type: orderType,
	}
	if 0 == len(query.order) {
		query.order = []orderBy{}
	}
	query.order = append(query.order, order)
	return query
}

func (query *QueryObject) cacheFoundItems(rows []interface{}) {
	query.cache = make(map[int]reflect.Value)
	for k, v := range rows {
		query.cache[k] = reflect.ValueOf(v)
	}
}

func (query QueryObject) sortFunction(orderType OrderType) func(x, y interface{}, dataType DataType) bool {
	switch orderType {
	case DESC:
		return func(x, y interface{}, dataType DataType) bool {
			return false == IndexResolvers.Less(dataType, x, y)
		}
	default:
		return func(x, y interface{}, dataType DataType) bool {
			return IndexResolvers.Less(dataType, x, y)
		}
	}
}

func (query *QueryObject) doOrderBy(rows []vdbRow) []vdbRow {
	less := IndexResolvers.Less
	sort.Slice(rows, func(i, j int) (res bool) {
		current := reflect.ValueOf(rows[i].Item())
		next := reflect.ValueOf(rows[j].Item())
		for _, by := range query.order {
			index := query.table.getIndex(by.Key)
			nextValue := next.FieldByName(by.Key).Interface()
			currentValue := current.FieldByName(by.Key).Interface()
			orderType := by.Type
			isEqual := false == less(index.DataType, nextValue, currentValue) && false == less(index.DataType, currentValue, nextValue)
			if false == isEqual {
				sorter := query.sortFunction(orderType)
				res = sorter(currentValue, nextValue, index.DataType)
				break
			}
		}
		return res
	})
	return rows
}

func (query *QueryObject) Delete() {
	_, pointers := query.filter()

	query.table.Delete(pointers...)
}
func (query *QueryObject) Limit(limit int) *QueryObject {
	query.limit = &limit
	return query
}
