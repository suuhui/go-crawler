package parser

import (
	"crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "寂寞成影萌宝")
	itemLen := 1
	if len(result.Items) != itemLen {
		t.Errorf("Result should have %d items. Got %d.", itemLen, len(result.Items))
	}
	var user model.UserProfile
	var ok bool
	if user, ok = result.Items[0].(model.UserProfile); !ok {
		t.Errorf("Items[0]'s type is not UserProfle")
	}
	expect := model.UserProfile{
		Name:      "寂寞成影萌宝",
		Age:       "83岁",
		Height:    "105CM",
		Income:    "财务自由",
		Marital:   "离异",
		Education: "初中",
		Job:       "金融",
		Home:      "南京市",
	}
	if user != expect {
		t.Errorf("error field. expect %v; got %v", expect, user)
	}
}
