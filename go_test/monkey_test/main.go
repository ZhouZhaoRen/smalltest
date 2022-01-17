package main

import (
	"fmt"
	"time"
)

func GetAge(a int64) int64 {
	return CalcAge(a) + a
}

func CalcAge(a int64) int64 {
	fmt.Println("调用CalcAge")
	if a == 0 {
		return 1
	}
	return 20
}

type User struct {
	Name     string
	Birthday string
}

// CalcAge 计算用户年龄
func (u *User) CalcAge(a int64) int {
	fmt.Println("调用")
	t, err := time.Parse("2006-01-02", u.Birthday)
	if err != nil {
		return -1
	}
	return int(time.Now().Sub(t).Hours()/24.0) / 365
}

// GetInfo 获取用户相关信息
func (u *User) GetInfo(a int64) string {
	fmt.Println(a)
	age := u.CalcAge(a)
	if age <= 0 {
		return fmt.Sprintf("%s很神秘，我们还不了解ta。", u.Name)
	}
	return fmt.Sprintf("%s今年%d岁了，ta是我们的朋友。", u.Name, age)
}
