package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type News struct {
	KeeperId primitive.ObjectID
	Title    string
	Content  bson.Raw
	DateTime time.Time
}
