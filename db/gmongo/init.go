package gmongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"utils/db/common"
)

type MgoClient struct {
	*mongo.Database
}

// NewMgoClient ...
func NewMgoClient(cfg *common.MongoConfig) (Cli *mongo.Database, err error) { //nolint
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.DBUri))
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return client.Database(cfg.DBDatabase), nil
}
