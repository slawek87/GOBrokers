package settings

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(
		Settings.Get("DATABASE_DRIVER"),
		Settings.Get("DATABASE_ARGS"),
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
