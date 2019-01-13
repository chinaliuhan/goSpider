package engine

//请求实例
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult //对下一个页面做解析
}

//解析用的实例
type ParseResult struct {
	Requests []Request
	Items    []interface{} //任意类型
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
