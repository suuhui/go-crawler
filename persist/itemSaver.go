package persist

import (
	"context"
	"crawler/engine"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemServer() (chan engine.Item, error) {
	out := make(chan engine.Item)
	//sniff用来给客户端维护集群状态
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	go func() {
		count := 0
		for {
			item := <-out
			count++
			log.Printf("Item server got #%d %T %+v\n", count, item, item)
			_ = save(client, item)
		}
	}()

	return out, nil
}

func save(client *elastic.Client, item engine.Item) (err error) {
	indexService := client.Index().BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err = indexService.Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
