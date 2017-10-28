package flatHouse

import (
	"github.com/slawek87/GOBrokers/elastic"
	"github.com/slawek87/GOBrokers/settings"
)

const INDEX_NAME = ".flathouse"

// Function indexes flat house data in elasticsearch.
// Default index for flat house documents is .flathouse-{Month}-{YYYY}
func IndexFlatHouses()  {
	var flatHouses []FlatHouseModel

	db, _ := settings.InitDB()
	defer db.Close()

	db.Where("elastic = ?", false).Find(&flatHouses)

	for count, _ := range flatHouses {
		db.Model(flatHouses[count]).Update("elastic", true)
		db.Model(&flatHouses[count]).Related(&flatHouses[count].Address, "Address")
        flatHouses[count].Elastic = true
		elastic.IndexDocument(flatHouses[count], INDEX_NAME, flatHouses[count].TypeOfBuildingStyle)
	}
}




