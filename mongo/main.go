package main

import (
	"context"
	"fmt"
	"github.com/pubgo/xerror"
	"github.com/pubgo/xprocess"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	// Majority read concern requested, but server was not started with --enableMajorityReadConcern. 28ms
	cc := c.Database("test_hello").
		Collection("test_hello", options.Collection().
			//SetReadPreference(readpref.PrimaryPreferred()).
			SetReadPreference(readpref.SecondaryPreferred()).
			SetWriteConcern(writeconcern.New(writeconcern.W(3)))) //writeconcern.WTimeout(2*time.Second)
	//SetWriteConcern(writeconcern.New(writeconcern.WMajority())).
	//SetReadConcern(readconcern.Majority()),

	//writeconcern.WTimeout(2*time.Second)

	//cc := c.Database("test_hello").Collection("test_hello")
	//cc1 := c.Database("test_hello").Collection("test_hello")
	var g = xprocess.NewGroup(100)
	defer g.Wait()
	for i := 0; ; i++ {
		i := i
		g.Go(func(ctx context.Context) {
			now := time.Now()
			dd, err := cc.InsertOne(context.Background(), bson.M{"foo": "bar", "hello": "world", "pi": 3.14159, "time": time.Now()})
			xerror.Panic(err)
			fmt.Println(dd.InsertedID.(primitive.ObjectID).Hex())
			//time.Sleep(time.Millisecond * 200)
			fmt.Println(time.Since(now))

			var data bson.M
			err = cc.FindOne(context.Background(), bson.M{"_id": dd.InsertedID}).Decode(&data)
			if err != nil {
				fmt.Println(err.Error(), time.Millisecond*time.Duration(i))
				return
			}
			fmt.Println(time.Since(now))
			fmt.Println(data)
			fmt.Println(g.Count())
		})
	}

	// 先从SecondaryPreferred读再从PrimaryPreferred读
}
