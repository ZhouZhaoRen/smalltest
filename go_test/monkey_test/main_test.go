package main

import (
	"fmt"
	"reflect"
	"testing"

	"bou.ke/monkey"
	c "github.com/smartystreets/goconvey/convey" // 别名导入
)

func TestGetAge(t *testing.T) {
	monkey.Patch(CalcAge, func(a int64) int64 {
		return 100
	})
	type args struct {
		a int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{name: "monkey_1", args: args{10}, want: 110},
		{name: "monkey_2", args: args{80}, want: 180},
	}
	for _, tt := range tests {
		c.Convey(tt.name, t, func() {
			if got := GetAge(tt.args.a); got != tt.want {
				t.Errorf("GetAge() = %v, want %v", got, tt.want)
			}
		})
		// 将t.Run换成convey
		//t.Run(tt.name, func(t *testing.T) {
		//
		//	if got := GetAge(tt.args.a); got != tt.want {
		//		t.Errorf("GetAge() = %v, want %v", got, tt.want)
		//	}
		//})
	}
}

func TestUser_GetInfo(t *testing.T) {
	type fields struct {
		Name     string
		Birthday string
	}
	type args struct {
		a int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{name: "getInfo_1", fields: fields{Name: "zzr", Birthday: "1995-03-08"}, args: args{a: 10}, want: fmt.Sprintf("%s今年%d岁了，ta是我们的朋友。", "zzr", 18)},
	}
	for _, tt := range tests {
		c.Convey(tt.name, t, func() {
			u := &User{
				Name:     tt.fields.Name,
				Birthday: tt.fields.Birthday,
			}
			// 为对象方法打桩
			monkey.PatchInstanceMethod(reflect.TypeOf(u), "CalcAge", func(*User, int64) int {
				return 18
			})

			if got := u.GetInfo(tt.args.a); got != tt.want {
				t.Errorf("GetInfo() = %v, want %v", got, tt.want)
			}
		})
		// // 将t.Run换成convey
		//t.Run(tt.name, func(t *testing.T) {
		//	u := &User{
		//		Name:     tt.fields.Name,
		//		Birthday: tt.fields.Birthday,
		//	}
		//	// 为对象方法打桩
		//	monkey.PatchInstanceMethod(reflect.TypeOf(u), "CalcAge", func(*User, int64) int {
		//		return 18
		//	})
		//
		//	if got := u.GetInfo(tt.args.a); got != tt.want {
		//		t.Errorf("GetInfo() = %v, want %v", got, tt.want)
		//	}
		//})
	}
}
