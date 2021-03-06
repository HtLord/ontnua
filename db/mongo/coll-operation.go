package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
TODO: Read some about dynamic type assign, reflect sort of stuffs to figure out. How to reduce looping assign from Find
*/

func CreateOne(coll *mongo.Collection, data interface{}) error {
	_, err := coll.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	return nil
}

func CreateMany(coll *mongo.Collection, data []interface{}) error {
	_, err := coll.InsertMany(context.TODO(), data)
	if err != nil {
		return err
	}
	return nil
}

func ReadOne(coll *mongo.Collection, filter interface{}, result interface{}) error {
	err := coll.FindOne(context.TODO(), filter).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func ReadMany(coll *mongo.Collection, filter interface{}) (*mongo.Cursor, error) {

	cur, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return cur, nil
}

/*
Use it with UpdateFilter. You can reference it by TestUpdateFilter in coll-operation_test.go
 */
func UpdateOne(coll *mongo.Collection, filter interface{}, data interface{}) error {
	_, err := coll.UpdateOne(context.TODO(), filter, data)
	if err != nil {
		return err
	}
	return nil
}

/*
When a request is be decode into json interface you can use this and assign as filter argument of UpdateOne without
declare a new variable and assign to it.
 */
func UpdateFilter(filter interface{}) interface{}{
	return bson.D{{"$set", filter}}
}