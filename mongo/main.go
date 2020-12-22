package main

import (
	"context"
	"fmt"
	"github.com/pubgo/xerror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"os"
	"time"
)

func main() {
	fmt.Println(os.Getenv("mongodb"))
	clientOptions := options.Client().ApplyURI(os.Getenv("mongodb"))
	c, err := mongo.NewClient(clientOptions)
	xerror.Panic(err)
	xerror.Panic(c.Connect(context.Background()))

	//cc := c.Database("test_hello").
	//	Collection("test_hello", options.Collection().SetReadPreference(readpref.SecondaryPreferred()))
	cc := c.Database("test_hello").
		Collection("test_hello", options.Collection().
			//SetReadPreference(readpref.PrimaryPreferred()).
			SetReadPreference(readpref.SecondaryPreferred()).
			SetWriteConcern(writeconcern.New(writeconcern.W(2))).
			SetReadConcern(readconcern.Majority()))

	//cc := c.Database("test_hello").Collection("test_hello")
	//cc1 := c.Database("test_hello").Collection("test_hello")
	for i := 0; ; i++ {
		dd, err := cc.InsertOne(context.Background(), bson.M{"foo": "bar", "hello": "world", "pi": 3.14159, "time": time.Now()})
		xerror.Panic(err)
		fmt.Println(dd.InsertedID.(primitive.ObjectID).Hex())
		time.Sleep(time.Millisecond * 200)

		var data bson.M
		err = cc.FindOne(context.Background(), bson.M{"_id": dd.InsertedID}).Decode(&data)
		if err != nil {
			fmt.Println(err.Error(), time.Millisecond*time.Duration(i))
			continue
		}
		//fmt.Println(data)
	}

	// 先从SecondaryPreferred读再从PrimaryPreferred读
}
