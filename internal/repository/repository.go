package repository

import (
	"database/sql"
	"os"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"github.com/stobita/plank/internal/usecase"
	"github.com/volatiletech/sqlboiler/boil"
)

type repository struct {
	db          *sql.DB
	esClient    *elastic.Client
	fileClient  *s3.S3
	redisClient *redis.Client
}

// New ...
func New(db *sql.DB, esClient *elastic.Client, fileClient *s3.S3, redisClient *redis.Client) usecase.Repository {
	if os.Getenv("PRODUCTION") != "true" {
		boil.DebugMode = true
	}
	return &repository{
		db:          db,
		esClient:    esClient,
		fileClient:  fileClient,
		redisClient: redisClient,
	}
}
