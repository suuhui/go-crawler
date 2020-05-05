package parser

import (
	"crawler/engine"
	"log"
	"regexp"
)

var profileRe = regexp.MustCompile(`<a href="(http://localhost:8080/mock/album.zhenai.com/u/\d+)">([^<]+)</a>`)
var cityRe = regexp.MustCompile(`href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[^"]+)"`)

func ParseCity(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}
	matches := profileRe.FindAllSubmatch(contents, -1)

	for _, match := range matches {
		if len(match) != 3 {
			log.Print("parse city error.", match)
		}
		name := string(match[2])
		url := string(match[1])
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(url, c, name)
			},
		})
	}

	cityMatches := cityRe.FindAllSubmatch(contents, -1)
	for _, match := range cityMatches {
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(match[1]),
			ParseFunc: ParseCity,
		})
	}

	return result
}
