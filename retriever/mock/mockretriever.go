package mock

import "fmt"

type Retriever struct {
	Content string
}

//35 、通过按 Ctrl-O （ Code | Override Methods… ）可以很容易地重载基本类地方法。
//
//要完成当前类 implements 的（或者抽象基本类的）接口的方法，就使用 Ctrl-I （ Code | Implement Methods… ）。

//就是看到fmt 中stringEr中的有一個String(),所以我們這裡,重寫一下試試,这个写完之后就相当于是其他语言中的toString
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
