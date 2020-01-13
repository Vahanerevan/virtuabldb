package virtuabldb

const TableName string = "User"

type User struct {
	Name   string
	Id     int
	Age    int
	IsMain bool
}

func createTable() {

	CreateTable(VirtualTable{
		Name: TableName,
		IndexOptions: IndexOptions{
			IndexOption{
				Name: "Id",
				Type: TYPE_INT,
			},
			IndexOption{
				Name: "Name",
				Type: TYPE_STRING,
			},
			IndexOption{
				Name: "Age",
				Type: TYPE_INT,
			},
		},
	})

}
