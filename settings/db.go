package settings

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	elastic "gopkg.in/olivere/elastic.v5"
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

func InitElasticSearch() *elastic.Client {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(
			Settings.Get("ELASTIC_SEARCH_URL")),
		elastic.SetBasicAuth(
			Settings.Get("ELASTIC_SEARCH_USERNAME"),
			Settings.Get("ELASTIC_SEARCH_PASSWORD")))

	if err != nil {
		panic(err)
	}

	return client
}