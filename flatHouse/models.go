package flatHouse

import (
	"github.com/jinzhu/gorm"
	"github.com/slawek87/GOBrokers/realEstate"
)

var BuildingMaterial = map[string]string {
	"brick": "brick",
	"concrete": "concrete",
	"wood":	"wood",
	"airBrick": "air brick",
	"differential": "differential",
	"stone": "stone",
	"blocks": "blocks",
	"frameH": "frame H",
	"suporex": "suporex"}

var TypeOfBuildingStyle = map[string]string {
	"apartment": "apartment",
	"flat": "flat",
	"multiFamilyHouse": "multi-family house",
	"tenementHouse": "tenement house"}


// Stores all information about flat houses added to system.
type FlatHouseModel struct {
	gorm.Model

	realEstate.RealEstateModel

	NumberOfRooms					int
	Floor							int
	NumberOfFloorsInBuilding		int
	YearOfConstruction				int
	BuildingMaterial				string

	TypeOfBuildingStyle				string		`gorm:"not null;"`
}
