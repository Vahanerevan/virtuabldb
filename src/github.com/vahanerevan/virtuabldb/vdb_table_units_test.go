package virtuabldb

import (
	"fmt"
	"testing"
)

func init() {
	createTable()
}

func BenchmarkTestVirtualTable_Insert_Benchmark(b *testing.B) {


	for i := 1; i <= b.N; i++ {
		strTest := fmt.Sprintf("test %v", i)

		table := Table(TableName)
		table.Insert(User{Name: strTest, Id: i, Age: i, IsMain: false})
	}
}

//func TestVirtualTable_Insert(test *testing.T) {
//	table := Table(TableName)
//
//	table.Insert(User{
//		Id:   1,
//		Name: "User",
//	})
//
//	table.Insert(User{
//		Id:   2,
//		Name: "User",
//	})
//	table.Insert(User{
//		Id:   3,
//		Name: "User",
//	})
//
//	if len(table.rows) != 3 {
//		test.Errorf("Insert crerate wrong amount of rows")
//	}
//
//	row := table.rows[0]
//
//	if row.ID != 0 {
//		test.Errorf("Wrong id generated %v", row.ID)
//	}
//}
//func TestQueryManager_Equals(test *testing.T) {
//	test.Logf("Filter Id Equal 1 %v", 1)
//	q := NewQuery(TableName)
//	result := q.Filter("Id").Equals(1).First()
//	if result.(User).Id != 1 {
//		test.Errorf("Error white filter equal id %v", result)
//	}
//}
//
//func TestQueryManager_NotEqual(test *testing.T) {
//	test.Logf("Filter Id Not Equal 1 %v", 1)
//	q := NewQuery(TableName)
//	result := q.Filter("Id").NotEqual(1).Result()
//	test.Log(result.Count())
//	result.Iterate(func(i interface{}) {
//		if i.(User).Id == 1 {
//			test.Errorf("Error white filter equal id %v", result)
//		}
//
//	})
//
//}
//func TestQueryManager_GreaterEqual(test *testing.T) {
//	q := NewQuery(TableName)
//	result := q.Filter("Id").GreaterEqual(1).Result()
//	test.Log(result.Length())
//	if result.First().(User).Id != 1 {
//		test.Errorf("Error white filter equal id %v", result)
//	}
//	if result.Last().(User).Id != 3 {
//		test.Errorf("Error white filter equal id %v", result)
//	}
//}
//
//func TestQueryManager_Greater(test *testing.T) {
//	q := NewQuery(TableName)
//	result := q.Filter("Id").Greater(1).Result()
//	if result.First().(User).Id != 3 {
//		test.Errorf("Error white filter equal id %v", result)
//	}
//	if result.Last().(User).Id != 2 {
//		test.Errorf("Error white filter equal id %v", result)
//	}
//}
