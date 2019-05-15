package back

import (
	"context"
	"github.com/HtLord/ontnua/model"
	"github.com/gin-gonic/gin"
	"time"
)

var dbName = "test1"
var collName = "keeper"

func CreateTest(ctx *gin.Context) {
	current := time.Now()

	name, ok := ctx.GetPostForm("name")
	acct, ok := ctx.GetPostForm("acct")
	secret, ok := ctx.GetPostForm("secret")
	if !ok {
		ctx.Error(nil)
		return
	}
	servmodel.
	_, err := GetColl(dbName, collName).InsertOne(context.TODO(), keeper)
	if err != nil{
		ctx.Error(nil)
		return
	}
}