package main

import (
	"container/list"
	"container/ring"
	"fmt"
)

func main() {
	test03()
}

// 测试环的用途，用于保存最近n条数据或者记录，自动剔除掉过期元素，保留最新的n条记录，
func test03() {
	r:=ring.New(5)
	for i:=1;i<20;i++ {
		r.Value=i
		r.Do(func(p interface{}) {
			fmt.Printf("%+v ",p)
		})
		fmt.Println()

		r=r.Next()
	}
}

// 测试环
func test02() {
	r := ring.New(10)
	for i := 1; i <= r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}
	r.Do(func(p interface{}) {
		fmt.Println(p)
	})
	//
	// 获得当前元素后面的第5个，当前r指向的是1，所以后面第5个是6
	r5 := r.Move(5)
	fmt.Println("r5==", r5.Value)
	fmt.Println("r==", r.Value)

	// 将当前元素连接到上面的元素，也就是去掉两个元素中间的元素，此时形成两个环，一个是链接后的环，另一个是去掉的元素形成的环，返回值是形成的环的第一个节点，也就是当前元素的下一个节点
	r1 := r.Link(r5)
	fmt.Println(r1.Value)
	fmt.Println(r.Value)
	r.Do(func(p interface{}) {
		fmt.Println(p)
	})
	fmt.Println("---------")
	r1.Do(func(p interface{}) {
		fmt.Println(p)
	})
}

// 测试list双向链表，双向链表的常规操作都支持，比如向前添加元素，往后添加元素，移动元素到某个位置
// 并且链表可以做为堆和栈的基础数据结构
func test01() {
	ls := list.New()
	for i := 97; i < 97+26; i++ {
		ls.PushFront(i)
	}
	// 遍历
	for it := ls.Front(); it != nil; it = it.Next() {
		fmt.Printf("%c ", it.Value)
	}
}
