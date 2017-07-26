package store

//import (
//	"testing"
//	"fmt"
//	"reflect"
//	//_ "github.com/go-sql-driver/mysql"
//	//"database/sql"
//)

//func TestSqlGenericStore_Run(t *testing.T) {
//	var Database = InitConnection()
//	mine := SqlGenericStore{Database}
//	res := <- mine.Select(map[string]interface{}{"id":3})
//	if res.Err != nil{
//		t.Errorf("test failed",res.Err)
//	}
//
//	fmt.Println(reflect.TypeOf(res.Data),res.Data)
//	//for val := range res.Data.(sql.Result).RowsAffected()
//}
