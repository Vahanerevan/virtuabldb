package vdatabase

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
				Type: TypeInt,
			},
			IndexOption{
				Name: "Name",
				Type: TypeString,
			},
			IndexOption{
				Name: "Age",
				Type: TypeInt,
			},
		},
	})

}
