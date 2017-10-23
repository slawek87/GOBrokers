package settings

import (
	elastic "gopkg.in/olivere/elastic.v5"
)

func InitElasticSearch() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(Settings.Get("ELASTIC_SEARCH_URL")))

	if err != nil {
		return nil, err
	}

	return client, nil
}
