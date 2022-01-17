package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("result==", test04())
}

func test04() bool {
	target := 31
	nums := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	if len(nums) == 0 || len(nums[0]) == 0 {
		return false
	}
	row, col := 0, len(nums[0])-1
	for row < len(nums) && col >= 0 {
		if target > nums[row][col] {
			row++
		} else if target < nums[row][col] {
			col--
		} else {
			return true
		}
	}

	return false
}


func test03() string {
	s := "we are happy"
	//
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			result = strings.Replace(s, " ", "%20", i)
			s = result
		}
	}

	ss := strings.Split(s, "")
	for i := 0; i < len(ss); i++ {
		if ss[i] == " " {
			ss[i] = "20%"
		}
	}
	//return strings.Join(ss,"")
	return result
}

func test02() string {
	s := "we are happy"
	ss := strings.Split(s, "")
	for i := 0; i < len(ss); i++ {
		if ss[i] == " " {
			ss[i] = "20%"
		}
	}
	return strings.Join(ss, "")
}

func test01() int {
	arr := []int{2, 3, 4, 1, 2, 0}
	for i := 0; i < len(arr); i++ {
		if arr[i] != i {
			cur := arr[i]
			if arr[cur] == cur {
				return cur
			} else {
				arr[i], arr[cur] = arr[cur], arr[i]
			}
		}
	}
	fmt.Println(arr)
	return 0
}
