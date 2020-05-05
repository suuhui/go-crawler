package persist

import (
	"context"
	"crawler/engine"
	"crawler/model"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"log"
	"reflect"
	"testing"
)

func TestItemServer(t *testing.T) {
	tests := []struct {
		name string
		want chan interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ItemServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ItemServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_save(t *testing.T) {
	profile := model.UserProfile{
		Name:      "寂寞成影萌宝",
		Age:       "83岁",
		Height:    "105CM",
		Income:    "财务自由",
		Marital:   "离异",
		Education: "初中",
		Job:       "金融",
		Home:      "南京市",
	}
	item := engine.Item{
		Url: "http://localhost:8080/mock/album.zhenai.com/u/8256018539338750764",
		Id: "8256018539338750764",
		Payload: profile,
	}
	err := save(item)
	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	result, err := client.Get().Index("user_profile").Id(item.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	user := model.UserProfile{}
	err = json.Unmarshal(result.Source, &user)
	if err != nil {
		panic(err)
	}

	if user != profile {
		t.Errorf("%+v, %+v", user, profile)
	}

	log.Printf("%v", user)
	log.Printf("%v", profile)
}
