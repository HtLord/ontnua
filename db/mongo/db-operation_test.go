package mongo

import (
	"context"
	"github.com/HtLord/ontnua/model"
	"github.com/stretchr/testify/assert"
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

func TestOneFilter(t *testing.T) {
	m := model.Member{}
	c := GetColl("test", "member")
	err := c.FindOne(context.TODO(), bson.D{{"name", "a1"}}).Decode(&m)
	if err != nil {
		t.Error(err)
		return
	}
	//
	//if err != nil {
	//	t.Error(err)
	//}

	assert.EqualValues(t, "a1", m.Name)
	assert.EqualValues(t, "b2", m.Email)
	assert.EqualValues(t, "c3", m.Secret)
}

func TestMultiFilter(t *testing.T) {
	m := model.Member{}
	c := GetColl("test", "member")
	cur, err := c.Find(context.TODO(), bson.D{{"name", "a1"}})
	if err != nil {
		t.Error(err)
		return
	}

	for cur.Next(context.TODO()) {
		err = cur.Decode(&m)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(m)
	}

}
