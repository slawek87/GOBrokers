package realEstate

import "github.com/slawek87/GOBrokers/settings"

func InitMigrations() {
	db, _ := settings.InitDB()
	defer db.Close()

	db.LogMode(true)
}
