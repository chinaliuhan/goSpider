package mock

type Retriever struct {
	Content string
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
