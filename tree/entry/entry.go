package main

import (
	"fmt"
	"learnGo/tree"
)

//通过组合, 扩展已有类型
type myTreeNode struct {
	//组合不一定要放一个指针,也可以放具体的类型
	node *tree.TreeNode
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	//这最好不要写成这样,虽然IDE没有报错, 但是编译的时候会提示地址不存在, 因为我们的myTreeNode是一个指针类型
	//myTreeNode{myNode.node.Left}.postOrder()
	//myTreeNode{myNode.node.Right}.postOrder()
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {

	//结构的创建
	var root tree.TreeNode

	root = tree.TreeNode{Value: 3}
	//这里就是将这个成员, 也设置成一个tradeNode的里面的内容
	//而且left和right是一个指针, 所以这里要通过传入指针类型的传入
	root.Left = &tree.TreeNode{}
	//将treeNode的right成员,设置成treeNode{5, nil, nil}的内容
	root.Right = &tree.TreeNode{5, nil, nil}
	//这里通过new 获取treeNode的地址,赋值给了root.right.left
	root.Right.Left = new(tree.TreeNode)

	//通过工厂方法来做赋值
	root.Left.Right = tree.CreateNode(2)

	//通过slice的形式来创建结构
	//nodes := []treeNode{
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)

	//这里我们要修改value的值实际上是不会生效的, 因为go语言所有的传值都是值类型传递
	// 所以我们需要把setValue的方法改成  func (node *treeNode) 即指针类型传递,然后这里在使用 root调用setValue的时候会被设置为值类型
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()

	//root.print()
	//root.setValue(100)

	//而go语言的所有参数都是值传递, 如果不是指针传递会拷贝一份传递给该方法
	//这里我们在实例化的时候获取到该对象的指针地址, 然后我们在调用setValude的时候,实际上这里会将这里实例化的地址传入到setValue的(node *treeNode)中
	//而将我们传递的值setValue(200)传递给setValude()&root传递给(node *treeNode)
	//也就是说, go语言的方法, 如果是一个值的话, 他会调用原来值来用
	//如果是一个指针的话, 他会调用原来的内存地址来用
	//这里会取刚才实例化后的root到内存地址进行调用
	//pRoot := &root
	//pRoot.print()
	//pRoot.setValue(200)
	//pRoot.print()

	//这里作为Nil指针的展示, go的nil指针也是可以传递的,但是注意在setValue的时候, 由于这个node并不存在,所以会报一个错误,但是传递时没问题的
	//而且下面的print也是一个nil
	var pRoot *tree.TreeNode ////这里会是一个Nil指针
	fmt.Println(pRoot)
	pRoot.SetValue(200)
	//可以看到里的调用和设置时没问题的,node不会存在==nil的确情况
	//而且下面的print也是有内容的
	pRoot = &root
	fmt.Println(pRoot)
	pRoot.SetValue(300)
	pRoot.Print()

	//数据,扩展类型的
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	//root.traverse()
}
