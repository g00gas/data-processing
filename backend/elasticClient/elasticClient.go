package elasticClient

import (
	"log"

	"github.com/olivere/elastic/v7"
)

var client *elastic.Client

func GetElasticsearchClient() (*elastic.Client, error) {
	var err error
	if client == nil {
		client, err = elastic.NewClient(elastic.SetURL("http://elasticsearch:9200"), elastic.SetBasicAuth("elastic", "password"), elastic.SetSniff(false))
		if err != nil {
			log.Fatalf("Failed to create Elasticsearch client: %v", err)
			return nil, err
		}
	}
	return client, nil
}
