package mysql_client

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/RandySteven/go-kopi/pkg/config"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

/*
### Repositories
<p>In repositories struct you can inject repository interface in Repositories</p>
*/
type (
	MySQL interface {
		Close()
		Ping() error
		Client() *sql.DB
		Migration(ctx context.Context) error
	}
	mysqlClient struct {
		db *sql.DB
	}
)

// Client implements [MySQL].
func (m *mysqlClient) Client() *sql.DB {
	return m.db
}

// Close implements [MySQL].
func (m *mysqlClient) Close() {
	m.db.Close()
}

// Migration implements [MySQL].
func (m *mysqlClient) Migration(ctx context.Context) error {
	return nil
}

// Ping implements [MySQL].
func (m *mysqlClient) Ping() error {
	return m.db.Ping()
}

func NewMYSQLClient(config *config.Config) (*mysqlClient, error) {
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
	return &mysqlClient{
		db: db,
	}, nil
}

var _ MySQL = &mysqlClient{}
