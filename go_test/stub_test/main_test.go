package main

import (
	"github.com/prashantv/gostub"
	"testing"
)

func TestFunc01(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
		{name: "test01", want: 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 对RPCFunc函数变量打桩，返回结果为1
			stubs := gostub.StubFunc(&RPCFunc, 1)
			// 对变量maxNum打桩，使得maxNum每次都是20
			stubs = stubs.Stub(&maxNum, 20)
			defer stubs.Reset()
			if got := Func01(); got != tt.want {
				t.Errorf("Func01() = %v, want %v", got, tt.want)
			}
		})
	}
}
