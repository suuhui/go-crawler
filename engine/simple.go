package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct {}

func (e *SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
	}
}

func worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetching error. url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	parseResult := r.ParseFunc(body)
	return parseResult, nil
}
