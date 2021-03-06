package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([^<]+)</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([^<]+)</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var maritalRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var jobRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var homeRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var idRe = regexp.MustCompile(`http://localhost:8080/mock/album.zhenai.com/u/(\d+)`)

func ParseProfile(url string, contents []byte, name string) engine.ParseResult {
	user := model.UserProfile{}
	user.Name = name
	user.Age = extractString(contents, ageRe)
	user.Education = extractString(contents, educationRe)
	user.Height = extractString(contents, heightRe)
	user.Income = extractString(contents, incomeRe)
	user.Marital = extractString(contents, maritalRe)
	user.Job = extractString(contents, jobRe)
	user.Home = extractString(contents, homeRe)

	matches := idRe.FindSubmatch([]byte(url))

	item := engine.Item{
		Url: url,
		Id: string(matches[1]),
		Payload: user,
	}

	result := engine.ParseResult{
		Items: []engine.Item{item},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
