package db

import (
	"database/sql"
	"github.com/RandySteven/go-kopi/interfaces/repositories"
)

/*
### Repositories
<p>In repositories struct you can inject repository interface in Repositories</p>
*/
type Repositories struct {
	/*
		####Dependencies injection
	*/
	UserRepository repositories.IUserRepository
	db             *sql.DB
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		db: db,
	}
}
