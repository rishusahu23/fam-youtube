package mongo

import (
	"context"
	"flag"
	"github.com/rishusahu23/fam-youtube/config"
	pkgMongo "github.com/rishusahu23/fam-youtube/pkg/test/mongo"
	mongo2 "github.com/rishusahu23/fam-youtube/user/dao/models/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	ctx := context.Background()
	conf, mongoClient, teardown := Initialise(ctx)
	defer teardown()
	UserDaoMongoTS = NewUserDaoMongo(mongoClient, conf)

	exitCode := m.Run()
	teardown()
	os.Exit(exitCode)
}

func Initialise(ctx context.Context) (*config.Config, *mongo.Client, func()) {
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}
	mongoClient := pkgMongo.SetUpTestDB(conf, mongo2.UserCollectionName)
	return conf, mongoClient, func() {
		_ = mongoClient.Disconnect(ctx)
	}
}
