package parser

import (
	"crawler/engine"
	"log"
	"regexp"
)

const cityRe = `<a href="(http://localhost:8080/mock/album.zhenai.com/u/\d+)">([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	result := engine.ParseResult{}
	matches := re.FindAllSubmatch(contents, -1)

	for _, match := range matches {
		if len(match) != 3 {
			log.Print("parse city error.", match)
		}
		name := string(match[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(match[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})

		result.Items = append(result.Items, name)
	}

	return result
}
