package flatHouse

import (
	"github.com/slawek87/GOBrokers/elastic"
	"github.com/slawek87/GOBrokers/settings"
)

const INDEX_NAME = "FlatHouse"

// Function indexes flat house data in elasticsearch.
// Default index for flat house documents is FlatHouse_{Month}.{YYYY}
func IndexFlatHouses()  {
	var flatHouses []FlatHouseModel

	db, _ := settings.InitDB()
	defer db.Close()

	db.Where("elastic = ?", false).Find(&flatHouses)

	for _, flatHouse := range flatHouses {
		db.Model(flatHouse).Update("elastic", true)
        flatHouse.Elastic = true
		elastic.IndexDocument(flatHouse, INDEX_NAME, flatHouse.TypeOfBuildingStyle)
	}
}




