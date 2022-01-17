package main

import "fmt"

type Greeting func( string) string

// 类型的方法
func (g Greeting) say(name string) {
	fmt.Println(g(name))
}

// 函数执行时将函数类型传进去，两种方法一样
func say(g Greeting, name string) {
	fmt.Println(g(name))
}

// 具体的实现函数
func english(name string) string {
	return "hello " + name
}

func main() {
	// 将具体的实现函数传进去运行
	say(english,"small")

	// 先实例化函数类型，再调用方法
	g:=Greeting(english)
	g.say("zhou")
	fmt.Println(g("hehe"))
}
