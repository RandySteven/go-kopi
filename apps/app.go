package apps

import (
	"github.com/RandySteven/go-kopi/pkg/config"
	mysql_client "github.com/RandySteven/go-kopi/pkg/db"
)

type (
	App struct {
		MySQL mysql_client.MySQL
	}
)

func NewApp(config *config.Config) (*App, error) {
	mysqlClient, err := mysql_client.NewMYSQLClient(config)
	if err != nil {
		return nil, err
	}
	return &App{
		MySQL: mysqlClient,
	}, nil
}
