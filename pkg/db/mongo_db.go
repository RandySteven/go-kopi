package db

import (
	"context"
	"fmt"
	"github.com/RandySteven/go-kopi/interfaces/repositories"
	"github.com/RandySteven/go-kopi/pkg/config"
	repositories2 "github.com/RandySteven/go-kopi/repositories"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepositories struct {
	DepositRepository repositories.IDepositRepository
	client            *mongo.Client
}

func NewMongoRepositories(config *config.Config) (*MongoRepositories, error) {
	url := fmt.Sprintf("mongodb://%s:%s@%s/%s", config.Mongodb.User, config.Mongodb.Password, config.Mongodb.Host, config.Mongodb.DbName)
	serversApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(url).SetServerAPIOptions(serversApi)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}
	return &MongoRepositories{
		DepositRepository: repositories2.NewDepositRepository(client),
		client:            client,
	}, nil
}
