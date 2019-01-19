package parser

import (
	"learnGo/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

//城市解析器,将城市的URL地址传进来,获取用户的URL地址
func ParseCity(contents []byte) engine.ParseResult {
	//生成正则表达式,一般我们自己写的用MustCompile,否则用Compile()处理错误信息
	re := regexp.MustCompile(cityRe)
	//返回一个[]byte ,相当于是一组被匹配到的字符串
	//matches := re.FindAll(contents, -1)
	//子匹配
	matches := re.FindAllSubmatch(contents, -1)
	//声明一个解析实例
	result := engine.ParseResult{}
	for _, m := range matches {
		//向实例中最追加城市名称
		result.Items = append(result.Items, "User "+string(m[2]))
		//请求实例中追加URL地址,同时拼装请求实例,追加到解析实例中的请求属性中
		name := m[2]
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			//这里通过返回一个匿名函数,匿名函数中调用ParseProfile进行解析个人信息,避免了ParserFunc还要加参数
			ParserFunc: func(content []byte) engine.ParseResult {
				//如果这里直接使用string(m[2]),会出现所有人的名字都是第一个人的问题
				//因为这里这个函数不是在这里运行,只是返回一个函数过去
				//真正运行是在循环完毕之后,在外面引擎中排队运行,等排到他之后,这里这个m早就不是那个人了
				//这个m的作用域,不是在for循环里面,而是所有这个for循环里面的返回的函数,所以我们把m[2]写到外面去重新定义一个变量copy一下
				return ParseProfile(content, string(name))
			},
		})
	}

	return result
}
