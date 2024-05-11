package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/go-kopi/entities/models"
	"github.com/RandySteven/go-kopi/interfaces/repositories"
	"github.com/RandySteven/go-kopi/utils"
)

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) Save(ctx context.Context, request *models.User) (result *uint64, err error) {
	return utils.Save[models.User](ctx, u.db, ``, request)
}

func (u *userRepository) FindAll(ctx context.Context) (result []*models.User, err error) {
	var user = &models.User{}
	result, err = utils.FindAll[models.User](ctx, u.db, ``, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userRepository) Find(ctx context.Context, id uint64) (result *models.User, err error) {
	result = &models.User{}
	err = utils.FindByID[models.User](ctx, u.db, ``, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userRepository) Delete(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) Update(ctx context.Context, request *models.User) (result *models.User, err error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

var _ repositories.IUserRepository = &userRepository{}
