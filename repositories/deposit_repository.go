package repositories

import (
	"context"
	"github.com/RandySteven/go-kopi/entities/models"
	"github.com/RandySteven/go-kopi/enums"
	"github.com/RandySteven/go-kopi/interfaces/repositories"
	"github.com/RandySteven/go-kopi/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type depositRepository struct {
	mongoClient *mongo.Client
}

func (d *depositRepository) Store(ctx context.Context, request *models.DepositData) (result *models.DepositData, err error) {
	result, err = utils.Store[models.DepositData](context.TODO(), d.mongoClient, enums.DepositData, request)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *depositRepository) FindAll(ctx context.Context) (result []*models.DepositData, err error) {
	//TODO implement me
	panic("implement me")
}

func (d *depositRepository) FindByID(ctx context.Context, id primitive.ObjectID) (result *models.DepositData, err error) {
	//TODO implement me
	panic("implement me")
}

func (d *depositRepository) Delete(ctx context.Context, id primitive.ObjectID) (err error) {
	//TODO implement me
	panic("implement me")
}

func (d *depositRepository) Update(ctx context.Context, request *models.DepositData) (result *models.DepositData, err error) {
	//TODO implement me
	panic("implement me")
}

func NewDepositRepository(mongoClient *mongo.Client) *depositRepository {
	return &depositRepository{
		mongoClient: mongoClient,
	}
}

var _ repositories.IDepositRepository = &depositRepository{}
