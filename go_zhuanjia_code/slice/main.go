package main

import (
	"fmt"
)

func main() {
	var array [10]int

	var slice = array[5:6]

	fmt.Println("lenth of slice: ", len(slice))
	fmt.Println("capacity of slice: ", cap(slice))
	fmt.Println(&slice[0] == &array[5])
}