package flatHouse

import (
	"github.com/slawek87/GOBrokers/settings"
	"context"
)

const INDEX_NAME = "FlatHouse"

// Function indexes flat house data in elasticsearch.
// Default index for flat house documents is FlatHouse_{Month}.{YYYY}
func IndexFlatHouse(flathouse *FlatHouseModel)  {
	client := settings.InitElasticSearch()
	_, err := client.Index().
	Index(settings.GetIndexName(INDEX_NAME)).
	BodyJson(flathouse).
	Do(context.Background())

	if err != nil {
		panic(err)
	}
}




