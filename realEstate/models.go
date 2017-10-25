package realEstate

var TypeOfTransaction = map[string]string{"sell": "sell", "rent": "rent"}
var PropertyType = map[string]string{"flatHouse": "flat house"}
var Currency = map[string]string{"pln": "PLN"}

// This is only Parent which should be inherited by child GORM structures.
// In this structure include basic RealEstate fields.
type RealEstateModel struct {
	Address     		AddressModel
	AddressID			int

	PropertyType		string       `gorm:"not null;"`

	TypeOfTransaction	string		 `gorm:"not null;"`
	Price				float32		 `gorm:"not null;"`
	Currency			string		 `gorm:"not null;"`

	Description			string		 `gorm:"not null;"`

	Elastic				bool		 `gorm:"not null;DEFAULT false;"`
}

// Stores address information.
type AddressModel struct {
	ID       		int

	City			string `gorm:"not null;"`
	Street			string
	District		string
}