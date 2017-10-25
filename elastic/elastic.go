package elastic

import (
	"strings"
	"time"
	"strconv"
	"github.com/slawek87/GOBrokers/settings"
	"context"
)

// function returns index name in pattern {name}_{Month}_{YYYY}.
func GetIndexName(name string) string {
	return strings.ToLower(name + "_" + time.Now().Month().String() + "_" + strconv.Itoa(time.Now().Year()))
}

// function indexes documents.
func IndexDocument(document interface{}, indexName string, documentType string)  {
	client := settings.InitElasticSearch()
	_, err := client.Index().
		Index(GetIndexName(indexName)).
		Type(documentType).
		BodyJson(document).
		Do(context.Background())

	if err != nil {
		panic(err)
	}
}
