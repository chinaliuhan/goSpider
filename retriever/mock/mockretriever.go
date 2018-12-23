package mock

type Retriever struct {
	Content string
	Title   string
}

//Retriever这里实现的接口和main.go中的无关,只要实现了GET方法方法就可以了
func (r Retriever) Get(url string) string {
	return r.Content
}
