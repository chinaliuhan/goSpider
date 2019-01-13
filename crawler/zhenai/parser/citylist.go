package parser

import (
	"learnGo/crawler/engine"
	"regexp"
)

const cityListRe = `<a\s+href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)"[^>]*>([^<]+)</a>`

//城市列表解析器,接收一段HTML,匹配URL列表,返回一个拼装好的解析实例
func ParseCityList(contents []byte) engine.ParseResult {
	//生成正则表达式,一般我们自己写的用MustCompile,否则用Compile()处理错误信息
	re := regexp.MustCompile(cityListRe)
	//返回一个[]byte ,相当于是一组被匹配到的字符串
	//matches := re.FindAll(contents, -1)
	//子匹配
	matches := re.FindAllSubmatch(contents, -1)
	//声明一个解析实例
	result := engine.ParseResult{}
	for _, m := range matches {
		//向实例中最追加城市名称
		result.Items = append(result.Items, string(m[2]))
		//请求实例中追加URL地址,同时拼装请求实例,追加到解析实例中的请求属性中
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}

	return result
}
