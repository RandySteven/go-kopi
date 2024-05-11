package db

import (
	"database/sql"
	"fmt"
	"github.com/RandySteven/go-kopi/interfaces/repositories"
	"github.com/RandySteven/go-kopi/pkg/config"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"log"
	"time"
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

func NewRepositories(config *config.Config) (*Repositories, error) {
	conn := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=require",
		config.Postgres.DbUser,
		config.Postgres.DbPass,
		config.Postgres.Host,
		config.Postgres.DbName,
	)
	log.Println(conn)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(8)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetConnMaxIdleTime(8 * time.Minute)
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return &Repositories{
		db: db,
	}, nil
}
