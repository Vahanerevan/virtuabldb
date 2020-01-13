package main

import (
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/shopspring/decimal"
	vdb "github.com/vahanerevan/virtuabldb"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"
)

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}

const TYPE_DECIMAL vdb.DataType = "decimal"
const TYPE_NULL_INT vdb.DataType = "null_int"

func init() {
	vdb.DefineIndexResolver(TYPE_DECIMAL, vdb.IndexResolver{
		Less: func(item, than interface{}) bool {
			return (item).(decimal.Decimal).LessThan(than.(decimal.Decimal))
		},
		Transform: nil,
	})

	vdb.DefineIndexResolver(TYPE_NULL_INT, vdb.IndexResolver{
		Less: func(item, than interface{}) bool {
			return *item.(*int) < *than.(*int)
		},
		Transform: func(item interface{}) interface{} {
			return item.(*int)
		},
	})
}

var schema vdb.Schema

func createTable() {

	vdb.CreateTable(vdb.VirtualTable{
		Name: "User",
		IndexOptions: vdb.IndexOptions{
			vdb.IndexOption{
				Name: "Id",
				Type: vdb.TYPE_INT,
			},
			vdb.IndexOption{
				Name: "Name",
				Type: vdb.TYPE_STRING,
			},
			vdb.IndexOption{
				Name: "Age",
				Type: vdb.TYPE_INT,
				Null: true,
			},
			vdb.IndexOption{
				Name: "IsMain",
				Type: vdb.TYPE_BOOL,
			},
			vdb.IndexOption{
				Name: "Balance",
				Type: TYPE_DECIMAL,
			},
		},
	})

	vdb.CreateTable(vdb.VirtualTable{
		Name: "UserPersonal",
		IndexOptions: vdb.IndexOptions{
			vdb.IndexOption{
				Name: "UserId",
				Type: vdb.TYPE_INT,
			},
			vdb.IndexOption{
				Name: "Address",
				Type: vdb.TYPE_STRING,
			},
			vdb.IndexOption{
				Name: "Phone",
				Type: vdb.TYPE_STRING,
			},
			vdb.IndexOption{
				Name: "Email",
				Type: vdb.TYPE_STRING,
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

	userTable := vdb.Table("User")
	personalTable := vdb.Table("UserPersonal")

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
	//wg.Add(2)
	//go func() {
	//	table.Insert(ins{Name: "username", Id: 1, Age: 32, IsMain: false, Balance: decimal.NewFromInt32(int32(3))})
	//	defer wg.Done()
	//}()
	//go func() {
	//	table.Insert(ins{Name: "username", Id: 2, Age: 12, IsMain: true, Balance: decimal.NewFromInt32(int32(5))})
	//	defer wg.Done()
	//}()
	//wg.Wait()
	insertStart()
	searchStart := timePoint("Search Time")

	q := vdb.NewQuery("User")
	q.Filter("Id").Equals(3)
	//q.Filter("Age").Equals( 5)
	result := q.Result()
	//pq.Where("Phone", "1")
	searchStart()

	//iup := pq.Filter("UserId").Equals(user{Id:4}.Id).First()
	//fmt.Println(iup.(userPersonal).UserId)
	fmt.Println(result.Count())
	result.Iterate(func(item interface{}) {
		fmt.Println(item)
		us := item.(user)
		fmt.Println("Found Item with Id", us.Id, " With Age", us.Age, "Balance", us.Balance)
	})

	totalStart()
	vdb.PrintMemUsage()

	var connection net.Conn
	http.ListenAndServe(":8798", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
			if err != nil {
			// handle error
			fmt.Println("ERRRRRRR",err)
		}
		fmt.Println("Client Connected")

		connection = conn

		go func() {

			//for {
			//
			//	reader := bufio
			//
			//	.NewReader(os.Stdin)
			//	fmt.Print("Enter text: ")
			//	text, _ := reader.ReadString('\n')
			//	fmt.Println(text)
			//	err = wsutil.WriteServerMessage(conn, ws.OpText, []byte("{}"))
			//}

			defer conn.Close()

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				fmt.Println("Message",string(msg))
				if err != nil {
					fmt.Println("ERRRRRRR",err)
					// handle error
				}
				err = wsutil.WriteServerMessage(conn, op, []byte("{}"))

				if err != nil {
					fmt.Println("ERRRRRRR",err)

					// handle error
				}
			}
		}()
	}))

}
