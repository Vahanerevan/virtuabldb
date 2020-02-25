package vdatabase

type QueryManager struct {
	key   string
	query *QueryObject
}

func (qm *QueryManager) Equals(value interface{}) *QueryObject {
	qm.query.addCondition(OpEqual, qm.key, value)
	return qm.query
}

func (qm *QueryManager) NotEqual(value interface{}) *QueryObject {
	qm.query.addCondition(OpNotEqual, qm.key, value)
	return qm.query
}

func (qm *QueryManager) IsNull() *QueryObject {
	qm.query.addCondition(OpIsNull, qm.key, nil)
	return qm.query
}

func (qm *QueryManager) IsNotNull() *QueryObject {
	qm.query.addCondition(OpIsNotNull, qm.key, nil)
	return qm.query
}
func (qm *QueryManager) In(values ...interface{}) *QueryObject {
	qm.query.addCondition(OpIn, qm.key, values)
	return qm.query
}

func (qm *QueryManager) Greater(value interface{}) *QueryObject {
	qm.query.addCondition(OpGt, qm.key, value)
	return qm.query
}

func (qm *QueryManager) GreaterEqual(value interface{}) *QueryObject {
	qm.query.addCondition(OpGte, qm.key, value)
	return qm.query
}

func (qm *QueryManager) Less(value interface{}) *QueryObject {
	qm.query.addCondition(OpLt, qm.key, value)
	return qm.query
}

func (qm *QueryManager) LessEqual(value interface{}) *QueryObject {
	qm.query.addCondition(OpLte, qm.key, value)
	return qm.query
}
