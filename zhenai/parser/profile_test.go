package parser

import (
	"crawler/engine"
	"crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile("http://localhost:8080/mock/album.zhenai.com/u/8256018539338750764", contents, "寂寞成影萌宝")
	itemLen := 1
	if len(result.Items) != itemLen {
		t.Errorf("Result should have %d items. Got %d.", itemLen, len(result.Items))
	}
	user := result.Items[0]
	payload := model.UserProfile{
		Name:      "寂寞成影萌宝",
		Age:       "83岁",
		Height:    "105CM",
		Income:    "财务自由",
		Marital:   "离异",
		Education: "初中",
		Job:       "金融",
		Home:      "南京市",
	}
	expect := engine.Item{
		Url:     "http://localhost:8080/mock/album.zhenai.com/u/8256018539338750764",
		Id:      "8256018539338750764",
		Payload: payload,
	}
	if user != expect {
		t.Errorf("error field. expect %v; got %v", expect, user)
	}
}
