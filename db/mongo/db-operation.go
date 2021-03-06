package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var _URI = os.Getenv("MONGO_SECRET")
var _Client *mongo.Client
var _DBPool []*mongo.Database
var _CollPool []*mongo.Collection

func getClient() *mongo.Client {
	if _Client == nil {
		if _URI == "" {
			panic("Can not found MONGO_SECRET in env_var.")
			return nil
		}
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(_URI))
		if err != nil {
			fmt.Println(err)
		}
		_Client = client
		fmt.Println("MongoDB client is created.")
		return _Client
	} else {
		return _Client
	}
}

func PingDB() {
	err := getClient().Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ping to MongoDB!")
	CloseDB()
}

func getDB(dbName string, opts ...*options.CollectionOptions) *mongo.Database {
	var db *mongo.Database
	for _, v := range _DBPool {
		if v.Name() == dbName {
			db = v
			break
		}
	}
	if db != nil {
		return db
	} else {
		db = getClient().Database(dbName)
		_DBPool = append(_DBPool, db)
		return db
	}
}

func GetColl(dbName string, collName string, opts ...*options.CollectionOptions) *mongo.Collection {
	var coll *mongo.Collection
	for _, v := range _CollPool {
		if v.Name() == collName {
			coll = v
			break
		}
	}
	if coll != nil {
		return coll
	} else {
		coll = getDB(dbName, opts...).Collection(collName)
		_CollPool = append(_CollPool, coll)
		return coll
	}
}

func CloseDB() int {
	if _Client == nil {
		fmt.Printf("No connection to MongoDB.\n")
		return 0
	}

	err := _Client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Connection to MongoDB closed.\n")

	return 1
}
