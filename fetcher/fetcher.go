package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

var defaultEncode = unicode.UTF8

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error status code: ", resp.StatusCode)
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bufioReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bufioReader)
	reader := transform.NewReader(bufioReader, e.NewDecoder())
	body, err := ioutil.ReadAll(reader)
	return body, nil
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	//DetermineEncoding会预读1024个byte确定编码
	//因此直接使用的话，后面会读不到前1024个字节
	//因此使用bufio.Peek读取前1024个字节不会使reader指针前进
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Print("determine encode fail, use default encode.")
		return defaultEncode
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
