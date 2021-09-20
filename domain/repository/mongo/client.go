package mongo

import (
	"context"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
	mongodb "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"massivleads/logger"
)

var (
	_            = godotenv.Load()
	client       = new(mongodb.Client)
	database     = os.Getenv("MONGO_DB_NAME")
	mongoOnce    = sync.Once{}
	mongoTimeout = getMongoTimeout()
)

func newMongoClient() *mongodb.Client {
	mongoOnce.Do(
		func() {

			ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
			defer cancel()
			c, err := mongodb.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URL")))

			if err != nil {
				logger.Error("Failed to init mongo client", err)
				panic("Fail to init mongo client")
			}
			err = c.Ping(ctx, readpref.Primary())
			if err != nil {
				logger.Error("Failed to ping mongo client", err)
				panic("Fail to ping mongo client")
			}

			client = c
		},
	)

	return client
}

func getMongoTimeout() time.Duration {
	tout, err := strconv.ParseInt(os.Getenv("MONGO_TX_TIMEOUT"), 10, 64)
	if err != nil {
		logger.Warn("Unable to parse mongo timeout using defaults", zap.Error(err))
		tout = 10
	}

	timeout := time.Duration(tout) * time.Second
	return timeout
}
