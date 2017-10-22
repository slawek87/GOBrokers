package realEstate


var TypeOfTransaction = map[string]string{"sell": "sell", "rent": "rent"}
var PropertyType = map[string]string{"flatHouse": "flatHouse"}

// Stores all basic information about realEstate. This model must be only inheritance.
type RealEstateModel struct {
	Address             AddressModel `gorm:"not null;unique"`

	PropertyType		string       `gorm:"not null;"`

	TypeOfTransaction	string		 `gorm:"not null;"`
	Price				float32		 `gorm:"not null;"`

	Description			string		 `gorm:"not null;"`
}

// Stores address information.
type AddressModel struct {
	City			string `gorm:"not null;"`
	Street			string
	District		string
}