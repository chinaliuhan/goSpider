package container

import "fmt"

//数组时值类型
func printArr(arr [5]int) {
	//这里也可以这么用, k值直接用 _ 占位符省略变量, 后面就可以不输出了下标, 值打印变量了
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}

//使用指针
func printArr1(arr *[5]int) {
	//这里也可以这么用, k值直接用 _ 占位符省略变量, 后面就可以不输出了下标, 值打印变量了
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	//数组,作为函数参数时是值类型传递,即函数会拷贝数组,函数里面的数组修改不会影响外面
	//如果想要作为类似于引用传递的效果,可以使用指针, go语言一般不直接使用数组, 而是用切片
	//数组的[5]int 和[10]int 在作为参数传递时, 是不同的类型
	//声明数组使用[], 可以指定数组的长足 数组类型依然放在后面 数量必须写在类型前面
	//使用var数组, 同时定义数组长度, 不给初始值的话, int的初始值即为0,比如这里输出5个0
	var arr1 [5]int
	//如果使用:=声明, 则必须有初始值
	arr2 := [3]int{1, 3, 5}
	//可以不声明长度,由编译器帮我们算长度, 但是必须有初始值
	arr3 := [...]int{2, 4, 6, 8, 10}

	//声明,4行5列的数组, 即二维数组: 里面有4个子数组,每个字数组中有5个值
	//这里如果用bool的话, 默认值为false
	var grid [4][5]bool

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	////使用for来, 遍历数组得到值 , 但是一般我们直接用range来做遍历
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	//使用range遍历数组,得到值
	for j := range arr3 {
		fmt.Println(arr3[j])
	}

	//通过遍历获取值和下标
	for k, v := range arr3 {
		fmt.Println(k, v)
	}

	printArr(arr1)
	//这里就会报错, 因为go语言中arr[5]和arr[3]是一个不同的类型,而printArr指定数组长度,必须传入相应长度的数组
	//printArr(arr2)
	//这里就不会报错,而且这里的0元素变成了100, 这是因为这里输出的函数里面的值
	printArr(arr3)
	//这里外面的输出,输出的结果依然是前面定义的值, 并没有被printArr里面的赋值所影响
	//可见数组确实是值传递
	fmt.Println("-----")
	fmt.Println(arr1, arr3)

	//这时候就会发现函数内部的修改会影响到外面
	printArr1(&arr1)
	fmt.Println("--111---")
	fmt.Println(arr1)
}
