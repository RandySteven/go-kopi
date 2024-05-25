package repositories

import "github.com/RandySteven/go-kopi/entities/models"

type IDepositRepository interface {
	MongoRepository[models.DepositData]
}
