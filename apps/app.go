package apps

import (
	"context"

	"github.com/RandySteven/go-kopi/caches"
	"github.com/RandySteven/go-kopi/handlers/consumers"
	handlers_rest "github.com/RandySteven/go-kopi/handlers/rest"
	"github.com/RandySteven/go-kopi/pkg/config"
	mysql_client "github.com/RandySteven/go-kopi/pkg/db"
	nsq_client "github.com/RandySteven/go-kopi/pkg/nsq"
	redis_client "github.com/RandySteven/go-kopi/pkg/redis"
	"github.com/RandySteven/go-kopi/repositories"
	"github.com/RandySteven/go-kopi/topics"
	"github.com/RandySteven/go-kopi/usecases"
)

type (
	App struct {
		MySQL mysql_client.MySQL
		Redis redis_client.Redis
		Nsq   nsq_client.Nsq
	}
)

func NewApp(config *config.Config) (*App, error) {
	mysqlClient, err := mysql_client.NewMYSQLClient(config)
	if err != nil {
		return nil, err
	}
	nsqClient, err := nsq_client.NewNsqClient(config)
	if err != nil {
		return nil, err
	}
	redisClient, err := redis_client.NewRedisClient(config)
	if err != nil {
		return nil, err
	}
	return &App{
		MySQL: mysqlClient,
		Redis: redisClient,
		Nsq:   nsqClient,
	}, nil
}

func (a *App) PrepareHttpHandler(ctx context.Context) *handlers_rest.Rests {
	repositories := repositories.NewRepositories(a.MySQL.Client())
	caches := caches.NewCaches(a.Redis.Client())
	usecases := usecases.NewUsecases(repositories, caches, &a.Nsq)
	return handlers_rest.NewRESTs(usecases)
}

func (a *App) RefreshRedis(ctx context.Context) error {
	return a.Redis.ClearCache(ctx)
}

func (a *App) PrepareJobScheduler(ctx context.Context) {
}

func (a *App) PrepareConsumer(ctx context.Context) *consumers.Consumers {
	topics := topics.NewTopics(a.Nsq)
	repositories := repositories.NewRepositories(a.MySQL.Client())
	caches := caches.NewCaches(a.Redis.Client())
	consumers := consumers.NewConsumers(repositories, caches, topics)
	return consumers
}
