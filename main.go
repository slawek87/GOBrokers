package main

import (
	"github.com/slawek87/GOBrokers/realEstate"
	"github.com/slawek87/GOBrokers/flatHouse"
	"github.com/slawek87/GOBrokers/settings"
	"github.com/slawek87/GOBrokers/elastic"
)

// Put here all methods and functions which you need to run before call main().
func init()  {
	realEstate.InitMigrations()
	flatHouse.InitMigrations()
}

func main()  {
    println("Starting!")
	db, _ := settings.InitDB()
	defer db.Close()

	address := realEstate.AddressModel{
		City: "Rybnik",
		District: "Sląskie",
		Street: "Żorska"}

	basicRealEstate := realEstate.RealEstateModel{
		PropertyType: realEstate.PropertyType["flatHouse"],

		TypeOfTransaction: realEstate.TypeOfTransaction["sell"],
		Currency: realEstate.Currency["pln"],

		Price: 500000,
		Description: "Testing description."}

	flatHouseRecord := flatHouse.FlatHouseModel{
		RealEstateModel: basicRealEstate,
		NumberOfRooms: 4,
		Floor: 5,
		NumberOfFloorsInBuilding: 5,
		YearOfConstruction: 1980,
		BuildingMaterial: flatHouse.BuildingMaterial["brick"],
		TypeOfBuildingStyle: flatHouse.TypeOfBuildingStyle["apartment"]}

    controller := flatHouse.FlatHouseController{}

    controller.CreateFlatHouseRecord(&flatHouseRecord, &address)

	flatHouse.IndexFlatHouses()

	elastic.IndexDocumentsTicker()
}
