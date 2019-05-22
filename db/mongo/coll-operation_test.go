package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

type TestObj struct {
	F1 string `bson:"f1" json:"f1"`
	F2 string `bson:"f2" json:"f2"`
}

func TestCreateOne(t *testing.T) {
	o := TestObj{"testF1", "testF2"}
	err := CreateOne(GetColl("test", "test"), o)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateMany(t *testing.T) {
	var o []interface{}

	o = append(o, TestObj{"t1", "t11"}, TestObj{"t2", "t11"}, TestObj{"t3", "t11"})
	err := CreateMany(GetColl("test", "test"), o)
	if err != nil {
		t.Error(err)
	}
}

func TestReadOne(t *testing.T) {
	f := bson.D{{"f1", "testF1"}}
	r := TestObj{}
	err := ReadOne(GetColl("test", "test"), f, &r)
	if err != nil {
		t.Error(err)
	}
	t.Log(r)
}

func TestReadMany(t *testing.T) {
	f := bson.D{{"f1", "t1"}}
	cur, err := ReadMany(GetColl("test", "test"), f)

	if err != nil {
		t.Error(err)
	}

	for cur.Next(context.TODO()) {
		var o TestObj
		bson.Unmarshal(cur.Current, &o)
		t.Log(o)
	}
}

func TestUpdateOne(t *testing.T) {
	f := bson.D{{"f1", "t1"}}
	u := bson.D{{"$set", bson.D{{"f2", "t22"},}},}

	err := UpdateOne(GetColl("test", "test"), f, u)
	if err != nil{
		t.Error(err)
	}
}

func TestUpdateFilter(t *testing.T) {
	f := bson.D{{"f1", "t1"}}
	u := bson.D{{"f2", "t22"}}
	err := UpdateOne(GetColl("test", "test"), f, UpdateFilter(u))
	if err != nil{
		t.Error(err)
	}
}