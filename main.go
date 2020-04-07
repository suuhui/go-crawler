package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

const startUrl = "http://localhost:8080/mock/www.zhenai.com/zhenghun"

func main() {
	engine.Run(engine.Request{
		Url:       startUrl,
		ParseFunc: parser.ParseCityList,
	})
}
