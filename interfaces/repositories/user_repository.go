package repositories

import "github.com/RandySteven/go-kopi/entities/models"

type IUserRepository interface {
	IRepository[models.User]
}
