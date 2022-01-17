package main

import (
	"fmt"
	"reflect"
)

func main() {
	test01()
}

func test01() {
	var x float64 = 3.4
	t := reflect.TypeOf(x)  //t is reflect.Type
	fmt.Println("type:", t)

	v := reflect.ValueOf(x) //v is reflect.Value
	fmt.Println("value:", v)
}
