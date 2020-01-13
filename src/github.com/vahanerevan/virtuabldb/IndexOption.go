package virtuabldb

type IndexOption struct {
	Name      string
	Type      DataType
	IsPrimary bool
	Null      bool
}
