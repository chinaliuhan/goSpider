package main

import (
	"testing"
)

//表格驱动测试

//可以用IDE左侧的箭头选择run进行测试
//也可以在命令行中,到当前目录下使用go test ./ 来测试

//代码测试的覆盖率通过函数IDE左侧的箭头run...converage 绿色为覆盖到的和覆盖次数,红色未没有覆盖
//也可以在当前目录下的命令行中 go test -coverprofile=c.out
// 然后使用go tool cover -html=c.out将结果转换为HTML 然后就会自动在浏览器中打开代码,绿色为覆盖到的和覆盖次数,红色未没有覆盖
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
		{"大大口香糖", 4}, //故意制造错误的话这里可以写0
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

//性能测试

//可以用IDE左侧的箭头选择run进行测试
//也可以在命令行中,到当前目录下使用go test -bench ./ 来测试
//输出结果BenchmarkSubstr-8   	 5000000代表输出多少次	       258 ns/op运行多少纳秒
//通过命令go test -bench . -cpuprofile cpu.out生成一个cpu.out二进制的文件,代码性能分析的文件.
//该文件通过go tool pprof ./cpu.out命令进入一个交互式的命令行,通过输入Help查看各种命令,输入web通过浏览器查看性能解析的页面
//通过help命令可以看到web Visualize graph through web browser 该命令可以通过浏览器查看性能分析页面
//注意 Failed to execute dot. Is Graphviz installed? Error: exec: "dot": executable file not found in $PATH
//在使用web命令前,需要安装Graphviz插件www.graphviz.org直接下载安装即可,mac版的直接用brew install Graphviz安装
//在分析页面中,方框越大,箭头越粗,说明运行时间越长
func BenchmarkSubstr(b *testing.B) {
	s := "大大口香糖"
	ans := 4
	//为了性能测试,这里把字符串搞的长一点
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s)=%d", len(s))
	//上面运行时间,不作为性能测试时间, 所以这里重置一下计时器
	b.ResetTimer()

	//既然是性能测试,一遍是不够的,所以我们这里进行循环
	//而循环的话,直接用b.N自带的算法结果,具体算多少遍不用我们去操心,系统会通过算法的记过告诉我们算多少遍,我们只要结果就好了
	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubstr1(s)
		if actual != ans {
			b.Errorf("got %d for input %s;"+"expected %d", actual, s, ans)
		}
	}
}
