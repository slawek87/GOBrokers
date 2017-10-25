package settings

import (
	elastic "gopkg.in/olivere/elastic.v5"
	"time"
	"strconv"
	"strings"
)

func InitElasticSearch() *elastic.Client {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(
			Settings.Get("ELASTIC_SEARCH_URL")),
		elastic.SetBasicAuth(
			Settings.Get("ELASTIC_SEARCH_USERNAME"),
			Settings.Get("ELASTIC_SEARCH_PASSWORD")))

	if err != nil {
		panic(err)
	}

	return client
}

// function returns index name in pattern {name}_{Month}_{YYYY}
func GetIndexName(name string) string {
	return strings.ToLower(name + "_" + time.Now().Month().String() + "_" + strconv.Itoa(time.Now().Year()))
}