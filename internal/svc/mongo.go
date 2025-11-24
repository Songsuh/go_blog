package svc

import (
	"github.com/Songsuh/go_blog/internal/global"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readconcern"
)

func CreatMongo(c *global.MongoDb) *mongo.Client {

	credential := options.Credential{
		AuthSource:    c.AuthSource,
		AuthMechanism: c.Mechanism,
		Username:      c.Username,
		Password:      c.Password,
	}
	opts := options.Client().
		ApplyURI(c.Uri).
		SetAuth(credential).
		SetReadConcern(readconcern.Majority()).
		SetMaxPoolSize(c.MaxPoolSize).
		SetMinPoolSize(c.MinPoolSize).
		SetTimeout(c.Timeout).
		SetConnectTimeout(c.ConnectTimeout).
		SetMaxConnIdleTime(c.MaxConnIdleTime)
	client, err := mongo.Connect(opts)

	if err != nil {
		panic(err)
	}

	return client
}
