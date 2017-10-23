package flatHouse

import (
	"github.com/slawek87/GOBrokers/realEstate"
	"github.com/slawek87/GOBrokers/settings"
)

type FlatHouseController struct{}

// method allows you to create FlatHouse record in DB with given Address data.
func (controller *FlatHouseController) CreateFlatHouseRecord(flatHouse *FlatHouseModel, address *realEstate.AddressModel) (*FlatHouseModel, error) {
	db, _ := settings.InitDB()
	defer db.Close()

	db.NewRecord(&address)
	query := db.Create(&address)

	if query.Error != nil {
		return flatHouse, query.Error
	}

	flatHouse.Address = *address
	flatHouse.AddressID = address.ID

	db.NewRecord(&flatHouse)
	query = db.Create(&flatHouse)

	return flatHouse, query.Error
}
