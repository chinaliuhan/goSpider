package main

import (
	"fmt"
)

type Human int

func (h Human) String() string {
	return fmt.Sprintf("%s", "String echo")
}

func (h Human) chinese(name string) (age int, sex string) {

	return 0, "nan"
}

//创建节点
type TradeNode struct {
	value       int
	left, right *TradeNode
}

//为节点的属性设置参数, 同时返回节点的内存地址
func CreateNode(value int) *TradeNode {
	return &TradeNode{value: value}
}

//打印节点
func (node TradeNode) print() {
	fmt.Println(node.value)
}

func (node *TradeNode) setValue(value int) {
	if node == nil {
		fmt.Println("setting value to nil")
		return
	}
	node.value = value
}

//遍历节点内容
func (node *TradeNode) traverse() {
	//当node 为nil 的时候跳过
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {

	//if err := os.Chmod("nihao.log", 777); err != nil {
	//	fmt.Println(err)
	//}
	//
	//var h Human
	//
	//fmt.Println(h.chinese("zemin"))

	var root TradeNode
	//将节点的value设置为3
	root = TradeNode{value: 3}
	//将节点内存地址赋值给节点的左子树
	root.left = &TradeNode{}
	//将节点内存地址赋值给节点的右子树,同时赋值
	root.right = &TradeNode{5, nil, nil}
	root.right.left = new(TradeNode)
	root.left.right = CreateNode(2)
	root.right.left.setValue(4)

	root.traverse()

	//传入一个空的指针,传入没问题,但是里面设置值的时候,会导致拿不到实例的属性而报错
	//var pRoot *TradeNode
	//pRoot.setValue(200)

}
