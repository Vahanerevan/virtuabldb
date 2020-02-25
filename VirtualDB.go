package vdatabase

type (
	IndexOptions []IndexOption
	RowPointer   = int
	IndexerType  int
	DataType     string
)

type OrderType string

const (
	ASC  OrderType = "asc"
	DESC           = "desc"
)


type pair struct {
	Key   string
	Value interface{}
}

var schema *Schema

func init() {
	schema = new(Schema)
	schema.tables = make(map[string]*VirtualTable)
}

type Schema struct {
	tables map[string]*VirtualTable
}

func Table(name string) *VirtualTable {
	return schema.Table(name)
}

func CreateTable(table VirtualTable) {
	table.Initialize()
	schema.addTable(&table)

}

func (schema *Schema) addTable(table *VirtualTable) {
	schema.tables[table.Name] = table
}

func (schema *Schema) getTable(tableName string) *VirtualTable {
	return schema.tables[tableName]
}

func (schema *Schema) Table(name string) *VirtualTable {
	return schema.getTable(name)
}
