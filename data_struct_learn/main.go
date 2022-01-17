package main

import "fmt"

func main() {
	//MapLearn()
	mapAndSlice()
}

func mapAndSlice() {
	s:=make([]int,0)  // 返回的是slice结构体，将这个结构体作为参数传到别的函数，不会影响当前，因为会重新复制一个slice结构体
	s=append(s,100)
	m:=make(map[string]string) // 返回的是*hmap指针，将这个指针作为参数传入别的函数，即使重新复制了一个，但新的map还是指向原来的地址，所以在新函数里面的修改会影响到之前的
	m["aa"]="bb"
	fmt.Printf("s==%+v   m==%+v\n",s,m)  // s==[100]   m==map[aa:bb]
	change(s,m)
	fmt.Printf("s==%+v   m==%+v\n",s,m) // s==[100]   m==map[aa:bb cc:dd]
}

func change(s []int,m map[string]string) {
	s=append(s,200)
	m["cc"]="dd"
}

func MapLearn() {
	//m1 := make(map[string]string)
	//m2 := make(map[string]string, 8)
}
