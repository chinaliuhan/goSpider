package main

import (
	"testing"
)

//表格驱动测试

//可以用IDE左侧的箭头进行测试
//也可以在命令行中,到当前目录下使用go test ./ 来测试
func TestSubstr(t *testing.T) {

	//这里是一个结构体数组
	//上面{}中的定义几个参数, 下面的{}中每一个{}都要补齐个数
	//而且这个特tests可以直接进行遍历,遍历的时候下标从0开始,值的子值用.取值
	tests := []struct {
		s   string
		ans int
	}{
		//normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},
		//Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbb", 1},
		{"abcabcabcd", 4},
		//chinese support
		{"这里是百度", 5},
		{"一二三三二一", 3},
		{"大大口香糖", 0}, //故意制造错误
	}
	//遍历结构体数组
	for _, tt := range tests {
		//将结构体数组中的属性,s传入到待测试的函数中
		actual := lengthOfNonRepeatingSubstr1(tt.s)
		//判断函数的返回值,是否等于预定好的结构体数组中的ans属性的值,不等于的话提示错误信息
		if actual != tt.ans {
			t.Errorf("got %d for input %s;"+"expected %d", actual, tt.s, tt.ans)
		}

	}
}
