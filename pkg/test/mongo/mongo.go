package mongo

import (
	"context"
	"github.com/rishusahu23/fam-youtube/config"
	pkgMongo "github.com/rishusahu23/fam-youtube/pkg/db/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpTestDB(conf *config.Config, collectionStr string) *mongo.Client {
	ctx := context.Background()
	client := pkgMongo.GetMongoClient(ctx, conf)
	collection := client.Database(conf.MongoConfig.MongoDBName).
		Collection(collectionStr)
	if err := collection.Drop(ctx); err != nil {
		panic(err)
	}
	return client
}
