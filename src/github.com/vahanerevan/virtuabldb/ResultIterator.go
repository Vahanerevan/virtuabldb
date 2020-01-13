package virtuabldb

type ResultIterator struct {
	cursor int
	items  []interface{}
}

func NewResultIterator(items []interface{}) *ResultIterator {
	return new(ResultIterator).SetItems(items)
}

func (resultIterator *ResultIterator) SetItems(_items []interface{}) *ResultIterator {

	resultIterator.items = _items
	resultIterator.Reset()
	return resultIterator
}

func (resultIterator *ResultIterator) Reset() {
	resultIterator.cursor = 0
}

func (resultIterator *ResultIterator) Next() bool {
	return resultIterator.has()
}

func (resultIterator *ResultIterator) Item() interface{} {
	result := resultIterator.get()
	resultIterator.do(1)
	return result
}

func (resultIterator *ResultIterator) do(value int) bool {
	resultIterator.cursor += value
	return true
}

func (resultIterator ResultIterator) has() bool {
	return resultIterator.hasIndex(resultIterator.cursor)
}

func (resultIterator ResultIterator) hasIndex(index int) bool {
	if resultIterator.cursor < 0 || index > len(resultIterator.items)-1 {
		return false
	}
	return true
}

func (resultIterator ResultIterator) get() interface{} {
	if false == resultIterator.has() {
		return nil
	}
	return resultIterator.items[resultIterator.cursor]
}

func (resultIterator *ResultIterator) Iterate(callable func(interface{})) {
	resultIterator.Reset()
	for resultIterator.Next() {
		callable(resultIterator.Item())
	}
}

func (resultIterator *ResultIterator) Last() interface{} {
	return resultIterator.items[len(resultIterator.items)-1]
}

func (resultIterator *ResultIterator) First() interface{} {
	return resultIterator.items[0]
}

func (resultIterator ResultIterator) Length() int {
	return len(resultIterator.items)
}

func (resultIterator ResultIterator) Count() int {
	return len(resultIterator.items)
}
