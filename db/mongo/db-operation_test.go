package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestGetColl(t *testing.T) {
	r, err := GetColl("test", "test").InsertOne(context.TODO(), bson.D{{"hello", "world"}})
	if err != nil {
		t.Error(err)
	}

	t.Log(*r)
}
