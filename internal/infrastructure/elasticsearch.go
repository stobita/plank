package infrastructure

import (
	"log"
	"os"

	"github.com/olivere/elastic/v7"
)

func NewESClient() (*elastic.Client, error) {
	return elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetURL(os.Getenv("ES_URL")),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
}
