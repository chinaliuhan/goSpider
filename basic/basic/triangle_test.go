package main

import "testing"

//表格驱动测试

//可以用IDE左侧的箭头选择run进行测试
//也可以在命令行中,到当前目录下使用go test ./ 来测试

//代码测试的覆盖率通过函数IDE左侧的箭头run...converage 绿色为覆盖到的和覆盖次数,红色未没有覆盖
//也可以在当前目录下的命令行中 go test -coverprofile=c.out
// 然后使用go tool cover -html=c.out将结果转换为HTML 然后就会自动在浏览器中打开代码,绿色为覆盖到的和覆盖次数,红色未没有覆盖
func TestTriangle(t *testing.T) {
	//这里是一个结构体数组
	//上面{}中的定义几个参数, 下面的{}中每一个{}都要补齐个数
	//而且这个特tests可以直接进行遍历,遍历的时候下标从0开始,值的子值用.取值
	tests := []struct{ a, b, c int }{
		{3, 4, 0},//这里故意写错一个观察运行结果
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}

	//遍历结构体数组,得到每一个{}中的值
	for _, tt := range tests {
		//将结构体数组中的属性,s传入到待测试的函数中,同时判断函数的返回值是否等于结构体数组中的属性中定义的属性的值
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d,%d);"+"got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}
