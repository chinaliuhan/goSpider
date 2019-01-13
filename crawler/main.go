package main

import (
	"learnGo/crawler/engine"
	"learnGo/crawler/zhenai/parser"
)

func main() {
	engine.Run(
		engine.Request{
			Url:        "http://www.zhenai.com/zhenghun",
			ParserFunc: parser.ParseCityList,
		},
	)

}
