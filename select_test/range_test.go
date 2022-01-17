package main

import (
	"fmt"
	"io"
	"os"
)

// 函数中使用for-range对切片进行遍历，获取切片的下标和元素素值，这里忽略函数的实际意义。
func RangeSlice(slice []int) {
	for index, value := range slice {
		_, _ = index, value
	}
	// 遍历过程中每次迭代会对index和value进行赋值，如果数据量大或者value类型为string时，对value的赋值操作可能是多余的，可以在for-range中忽略value值，使用slice[index]引用value值。
}

//  函数中使用for-range对map进行遍历，获取map的key值，并根据key值获取获取value值，这里忽略函数的实际意义。
func RangeMap(myMap map[int]string) {
	for key, _ := range myMap {
		_, _ = key, myMap[key]
	}
	// 函数中for-range语句中只获取key值，然后根据key值获取value值，虽然看似减少了一次赋值，但通过key值查找value值的性能消耗可能高于赋值消耗。能否优化取决于map所存储数据结构特征、结合实际情况进行。
}

//  main()函数中定义一个切片v，通过range遍历v，遍历过程中不断向v中添加新的元素。
func main() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	// 能够正常结束。循环内改变切片的长度，不影响循环次数，循环次数在循环开始前就已经确定了。
}

func testReflect() {
	var r io.Reader
	tty, err := os.OpenFile("", os.O_RDWR, 0)
	if err != nil {
		return
	}
	r = tty
	fmt.Println(r)
}
