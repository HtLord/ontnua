package api

import (
	"context"
	"fmt"
	"github.com/HtLord/ontnua/db/mongo"
	"github.com/HtLord/ontnua/model"
	"github.com/gin-gonic/gin"
	"time"
)

var coll = mongo.GetColl(MemberDocNameMap.DbName, MemberDocNameMap.CollName)

func MemberRouting(r *gin.Engine){
	r.POST("/create/member", CreateMember)
	r.GET("/update/member/:mail/:phoneNumber", CreateMember)
	r.POST("/update/member", CreateMember)
	r.GET("/read/member/:mail/:phoneNumber", ReadMember)	
}

func CreateMember(c *gin.Context) {

	m := model.Member{}
	err := c.ShouldBind(&m)
	if err != nil{
		fmt.Println(err)
		return
	}

	current := time.Now()
	m.CreateTime = current
	m.UpdateTime = current

	_, err = coll.InsertOne(context.TODO(), m)
	if err != nil{
		fmt.Println(err)
		return
	}
}

func UpdateMember(c *gin.Context) {

	m := model.Member{}
	err := c.ShouldBind(&m)
	if err != nil{
		fmt.Println(err)
		return
	}

	current := time.Now()
	m.CreateTime = current
	m.UpdateTime = current

	_, err = coll.InsertOne(context.TODO(), m)
	if err != nil{
		fmt.Println(err)
		return
	}
}

func ReadMember(c *gin.Context) {

	m := model.Member{}
	err := c.ShouldBind(&m)
	if err != nil{
		fmt.Println(err)
		return
	}

	current := time.Now()
	m.CreateTime = current
	m.UpdateTime = current

	_, err = coll.InsertOne(context.TODO(), m)
	if err != nil{
		fmt.Println(err)
		return
	}
}

func DeleteMember(c *gin.Context) {

	m := model.Member{}
	err := c.ShouldBind(&m)
	if err != nil{
		fmt.Println(err)
		return
	}

	current := time.Now()
	m.CreateTime = current
	m.UpdateTime = current

	_, err = coll.InsertOne(context.TODO(), m)
	if err != nil{
		fmt.Println(err)
		return
	}
}