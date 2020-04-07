package main

import (
	"fmt"
	"regexp"
)

const text = `
email is test@qq.com
email2 is hello@163.com
email3 is world@gmail.com
email4 is golang@abc.com.cn
`

func main() {
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+[a-zA-z0-9]+`)
	match := re.FindAllString(text, -1) //-1表示找所有
	fmt.Println(match)

	//子匹配
	subRe := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-z0-9.]+)`)
	matches := subRe.FindAllStringSubmatch(text,  -1)
	for _, m := range matches {
		fmt.Println(m)
	}
}
