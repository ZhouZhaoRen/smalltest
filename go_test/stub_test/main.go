package main

import (
	"encoding/json"
	"fmt"
)

var maxNum int

type Person struct {
	Id   int
	name string
}

func main() {

}

func Func01() int {
	res, _ := json.Marshal(Person{
		Id: 1, name: "string",
	})
	var person Person
	err := json.Unmarshal(res, &person)
	if err != nil {
		fmt.Println(err)
	}

	return RPCFunc()+maxNum
}

var RPCFunc = func() int {
	fmt.Println("调用rpc")
	return maxNum
}
