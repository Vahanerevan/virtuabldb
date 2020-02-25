package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/vahanerevan/virtuabldb"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}

const TypeDecimal vdatabase.DataType = "decimal"
const TypeNullInt vdatabase.DataType = "null_int"

func init() {
	vdatabase.DefineIndexResolver(TypeDecimal, vdatabase.IndexResolver{
		Less: func(item, than interface{}) bool {
			return (item).(decimal.Decimal).LessThan(than.(decimal.Decimal))
		},
		Transform: nil,
	})

	vdatabase.DefineIndexResolver(TypeNullInt, vdatabase.IndexResolver{
		Less: func(item, than interface{}) bool {
			return *item.(*int) < *than.(*int)
		},
		Transform: func(item interface{}) interface{} {
			return item.(*int)
		},
	})
}

func createTable() {

	vdatabase.CreateTable(vdatabase.VirtualTable{
		Name: "User",
		IndexOptions: vdatabase.IndexOptions{
			vdatabase.IndexOption{
				Name: "Id",
				Type: vdatabase.TypeInt,
			},
			vdatabase.IndexOption{
				Name: "Name",
				Type: vdatabase.TypeString,
			},
			vdatabase.IndexOption{
				Name: "Age",
				Type: vdatabase.TypeInt,
				Null: true,
			},
			vdatabase.IndexOption{
				Name: "IsMain",
				Type: vdatabase.TypeBool,
			},
			vdatabase.IndexOption{
				Name: "Balance",
				Type: TypeDecimal,
			},
		},
	})

	vdatabase.CreateTable(vdatabase.VirtualTable{
		Name: "UserPersonal",
		IndexOptions: vdatabase.IndexOptions{
			vdatabase.IndexOption{
				Name: "UserId",
				Type: vdatabase.TypeInt,
			},
			vdatabase.IndexOption{
				Name: "Address",
				Type: vdatabase.TypeString,
			},
			vdatabase.IndexOption{
				Name: "Phone",
				Type: vdatabase.TypeString,
			},
			vdatabase.IndexOption{
				Name: "Email",
				Type: vdatabase.TypeString,
			},
		},
	})

}

func durationFormat(duration time.Duration) string {
	mis := duration.Microseconds()
	return fmt.Sprintf("%f seconds", float64(mis)/(1000*1000))
}

var max = 51

type user struct {
	Name    string
	Id      int
	Age     *int
	IsMain  bool
	Balance decimal.Decimal
}

type userPersonal struct {
	UserId  int
	Address string
	Phone   string
	Email   string
}

func timePoint(message string) func() {
	now := time.Now()
	return func() {
		elapsed := time.Since(now)
		df := durationFormat(elapsed)
		fmt.Println(fmt.Sprintf(message+"  %s  ", df))
	}
}

func main() {
	//var wg sync.WaitGroup
	createTable()
	rand.Seed(time.Now().UnixNano())
	totalStart := timePoint("Total Time")

	userTable := vdatabase.Table("User")
	personalTable := vdatabase.Table("UserPersonal")

	insertStart := timePoint("Insert Time")

	for i := 1; i <= max; i++ {
		strTest := fmt.Sprintf("test %v", i)
		var age *int = nil
		if i%2 != 0 {
			age = &i
		}
		userTable.Insert(user{Name: strTest, Id: i, Age: age, IsMain: false, Balance: decimal.NewFromInt32(int32(i))})
		personalTable.Insert(userPersonal{UserId: i, Address: "my address " + strconv.Itoa(i), Phone: strconv.Itoa(i)})
	}

	insertStart()
	searchStart := timePoint("Search Time")

	q := vdatabase.Query("User")
	q.Filter("Id").Equals(3)
	result := q.Result()

	searchStart()

	fmt.Println(result.Count())
	result.Iterate(func(item interface{}) {
		fmt.Println(item)
		us := item.(user)
		fmt.Println("Found Item with Id", us.Id, " With Age", us.Age, "Balance", us.Balance)
	})

	totalStart()
	vdatabase.PrintMemUsage()


}
