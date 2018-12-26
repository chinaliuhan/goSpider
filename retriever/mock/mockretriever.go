package mock

import "fmt"

type Retriever struct {
	Content string
}

//35 、通过按 Ctrl-O （ Code | Override Methods… ）可以很容易地重载基本类地方法。
//
//要完成当前类 implements 的（或者抽象基本类的）接口的方法，就使用 Ctrl-I （ Code | Implement Methods… ）。

//就是看到fmt 中stringEr中的有一個String(),所以我們這裡,重寫一下試試,这个写完之后就相当于是其他语言中的toString
//当采用任何接受字符的verb（%v %s %q %x %X）动作格式化一个操作数时，或者被不使用格式字符串如Print函数打印操作数时，会调用String方法来生成输出的文本。
//其他的实例的输出因为没有实现String()所以输出的都是默认的系统输出
func (r *Retriever) String() string {
	return fmt.Sprintf("!!!Retriever:{Contests=%s}", r.Content)
}

//这里要修改属性的值,必须用指针方式
func (r *Retriever) Post(url string, form map[string]string) string {
	r.Content = form["Contents"]
	return "ok"
}

//Retriever这里实现的接口和main.go中的无关,只要实现了GET方法方法就可以了
func (r Retriever) Get(url string) string {
	return r.Content
}
