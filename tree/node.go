package tree

import "fmt"

//封装方式
//名字一般采用CamelCase 作为
//首字母大写 相当于 public
//首字母小写 相当于private

//包
//包名和目录名不一定要一样, 但是每个目录只能有一个包
//main包包含可执行入口, 里面有一个main函数, 所以我们这个目录下有一个main函数, 那这里只能有一个main包否则我们可以给目录起其他的名字
//为结构定义的方法必须定义在同一个包内, 但是截图构体可以是不同的文件,比如有些

//扩展已有类型,系统类型或者别人的类型
//定义别名 或者 使用组合

//面向对象编程
//go语言的仅支持封装,不支持继承和多态,所以没有class,只有struct(结构体)
//go语言的面向对象, 没有继承和多态 只有封装
//只有面向接口编程
//这里着重介绍, 定义方法时的两种方式和 方法作为值接受者和指针接受者时的不同
//我们的编译器在调的时候, 他会自动识别这个函数是要值还是要指针要值他会取值, 要指针的他会自动取出内存地址
//不论是地址还是结构本身一律使用.来访问成员
//要改变内容的话,必须使用指针接受者,值接收者他只是一个拷贝
//因为值接收者是一个拷贝, 如果这个数据过大,要考虑性能问题,建议使用指针接受者
//一致性, 如果有指针接受者,最好就用指针接受者
//值接受者是go语言的特有
///值/指针 接收者均可以接收值/指针

type TreeNode struct {
	Value       int       //这样的属性,默认是0  键是Value值是0 int类型
	Left, Right *TreeNode //这样的属性,默认是nil 键是left和right 值是nil 类型是指针类型TreeNode
}

//为treeNode创建一个方法
//这里的(node tradeNode)是接收者的意思 而 print()是当前的方法名
//这个函数的意思也就是说, 这个Print()treeNode()接收的,也就是treeNode的方法
//这样写的话,调用的时候就需要用 treeNode.print 也可以将下面的方法写作print(node treeNode) 但是调用的时候就需要用print()这样来调
func (node TreeNode) Print() {
	fmt.Println(node.Value)
}

//为treeNode新增一个设置value的方法
//注意go语言的所有的参数传递都是值传递, 这里设置完value之后再调用上面的print()输出的话, 还会是原来的值
//如果要设置修改的话, 需要将调用 treeNode 的形式 定义为指针类型调用,  也就是*treeNode 否则默认就是值类型的调用
//注意这里的(node *treeNode) 也是类似于传入进来的, 下面的最后一次调用会有简单说明
//指针接收者,不加*为值接收者
func (node *TreeNode) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node,Ignore")
		return
	}
	node.Value = value
}

//go语言没有构造函数, 如果需要构造函数,可以通过这样的工厂函数来做,以实现构造函数的方式
//工厂函数,返回一个结构地址即可
func CreateNode(value int) *TreeNode {
	//在c++里面这里返回的是一个局部变量的地址注意这里的&符,c++里面会挂掉, 但是go语言不会
	return &TreeNode{Value: value}
}
