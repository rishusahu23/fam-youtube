//go:build wireinject
// +build wireinject

package wire

import (
	"crypto/tls"
	"github.com/google/wire"
	redisV9 "github.com/redis/go-redis/v9"
	"github.com/rishusahu23/fam-youtube/config"
	"github.com/rishusahu23/fam-youtube/external/post"
	"github.com/rishusahu23/fam-youtube/pkg/in_memory_store/redis"
	mongo2 "github.com/rishusahu23/fam-youtube/pkg/transaction/mongo"
	"github.com/rishusahu23/fam-youtube/user"
	mongoDao "github.com/rishusahu23/fam-youtube/user/dao/mongo"
	strategy "github.com/rishusahu23/fam-youtube/user/getuserstrategy"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func GetPostClientProvider(provider *post.ClientImpl) post.Client {
	return provider
}

func getHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}

func InitialiseUserService(conf *config.Config, mongoClient *mongo.Client, redisClient *redisV9.Client) *user.Service {
	wire.Build(
		user.NewService,
		mongoDao.UserDaoWireSet,
		mongo2.DefaultTxnManagerWireSet,
		strategy.FactoryWireSet,
		strategy.NewDB,
		strategy.NewCache,
		redis.RedisWireSet,
	)
	return &user.Service{}
}
