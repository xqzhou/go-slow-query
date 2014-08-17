package main

import (
	"fmt"
//	"strconv"
	"reflect"
	"github.com/youtube/vitess/go/vt/sqlparser"
)

//type Vertex struct {
//	X int
//	Y int
//}


func main() {
	statement, err := sqlparser.Parse("select * from a where b = 1")
	if err != nil {
		fmt.Println(err)
	}
	v := reflect.ValueOf(statement)
	fmt.Println(v)
	t := reflect.TypeOf(statement)
	fmt.Println(t)
//	fmt.Println(statement.Where)
//	fmt.Println()
	sel := statement.(*sqlparser.Select)
	fmt.Println(sel.Where.Type)
	expr := sel.Where.Expr.(*sqlparser.ComparisonExpr)
	fmt.Println(expr.Operator)
	fmt.Println(string(expr.Left.(*sqlparser.ColName).Name))
//	fmt.Println(strconv.Atoi(expr.Right.(sqlparser.NumVal)[0]))

//	v := Vertex{1, 2}
//	fmt.Println(v)
}
