package flatHouse

import (
	"github.com/slawek87/GOBrokers/elastic"
)

const INDEX_NAME = "FlatHouse"

// Function indexes flat house data in elasticsearch.
// Default index for flat house documents is FlatHouse_{Month}.{YYYY}
func IndexFlatHouse(flathouse *FlatHouseModel)  {
	elastic.IndexDocument(flathouse, INDEX_NAME, flathouse.TypeOfBuildingStyle)
}




