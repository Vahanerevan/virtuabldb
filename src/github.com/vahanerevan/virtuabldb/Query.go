package virtuabldb

import (
	"github.com/juliangruber/go-intersect"
	"github.com/spf13/cast"
	"reflect"
	"sort"
)

func NewQuery(tableName string) *Query {
	return &Query{
		table:      Table(tableName),
		conditions: map[OperationType][]pair{},
	}
}

func resetQuery(q *Query) {
	q.conditions = map[OperationType][]pair{}
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

type QueryManager struct {
	key   string
	query *Query
}

func (qm *QueryManager) Equals(value interface{}) *Query {
	qm.query.addCondition(OP_EQUAL, qm.key, value)
	return qm.query
}

func (qm *QueryManager) NotEqual(value interface{}) *Query {
	qm.query.addCondition(OP_NOT_EQUAL, qm.key, value)
	return qm.query
}

func (qm *QueryManager) IsNull() *Query {
	qm.query.addCondition(OP_IS_NULL, qm.key, nil)
	return qm.query
}

func (qm *QueryManager) IsNotNull() *Query {
	qm.query.addCondition(OP_IS_NOT_NULL, qm.key, nil)
	return qm.query
}
func (qm *QueryManager) In(values ...interface{}) *Query {
	qm.query.addCondition(OP_IN, qm.key, values)
	return qm.query
}

func (qm *QueryManager) Greater(value interface{}) *Query {
	qm.query.addCondition(OP_GT, qm.key, value)
	return qm.query
}

func (qm *QueryManager) GreaterEqual(value interface{}) *Query {
	qm.query.addCondition(OP_GTE, qm.key, value)
	return qm.query
}

func (qm *QueryManager) Less(value interface{}) *Query {
	qm.query.addCondition(OP_LT, qm.key, value)
	return qm.query
}

func (qm *QueryManager) LessEqual(value interface{}) *Query {
	qm.query.addCondition(OP_LTE, qm.key, value)
	return qm.query
}

type Query struct {
	table      *VirtualTable
	conditions map[OperationType][]pair
	cache      map[int]reflect.Value
	order      []orderBy
	limit      *int
}

func (query *Query) TableRowCount() int {
	return query.table.Count()
}

func (query *Query) Index(name string) *QueryManager {
	return &QueryManager{
		key:   name,
		query: query,
	}
}

func (query *Query) Filter(name string) *QueryManager {
	return query.Index(name)
}

func (query *Query) Where(name string, value interface{}) *Query {
	query.addCondition(OP_EQUAL, name, value)
	return query
}

func (query *Query) addCondition(operationType OperationType, key string, value interface{}) {
	query.conditions[operationType] = append(query.conditions[operationType], pair{
		Key:   key,
		Value: value,
	})
}

func (query *Query) filter(isSingle bool) (rows []interface{}, pointers []int) {

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

	vdbRows := query.table.CollectRowPointers(pointers)

	//query.cacheFoundItems(vdbRows)

	if 0 != len(query.order) {
		vdbRows = query.doOrderBy(vdbRows)
	}

	rows = collectRowsData(vdbRows)
	return rows, pointers
}

func (query *Query) Result() *ResultIterator {
	rows, _ := query.filter(false)
	resetQuery(query)
	return NewResultIterator(rows)
}
func (query *Query) First() interface{} {
	rows, _ := query.filter(true)
	resetQuery(query)
	if len(rows) > 0 {
		return rows[0]
	}

	return nil
}

func (query *Query) OrderBy(name string, orderType OrderType) *Query {
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

func (query *Query) cacheFoundItems(rows []interface{}) {
	query.cache = make(map[int]reflect.Value)
	for k, v := range rows {
		query.cache[k] = reflect.ValueOf(v)
	}
}

func (query Query) createSortFunction(orderType OrderType) func(x, y interface{}, dataType DataType) bool {
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

func (query *Query) doOrderBy(rows []vdbRow) []vdbRow {
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
				sorter := query.createSortFunction(orderType)
				res = sorter(currentValue, nextValue, index.DataType)
				break
			}
		}
		return res
	})
	return rows
}

func (query *Query) Delete() {
	_, pointers := query.filter(false)

	query.table.Delete(pointers...)
}
func (query *Query) Limit(limit int) *Query {
	query.limit = &limit
	return query
}
