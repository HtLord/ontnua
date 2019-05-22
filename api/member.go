package api

import (
	"context"
	"fmt"
	"github.com/HtLord/ontnua/db/mongo"
	"github.com/HtLord/ontnua/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var coll = mongo.GetColl(MemberDocNameMap.DbName, MemberDocNameMap.CollName)

func MemberRouting(r *gin.Engine) {
	r.POST("/create/member", CreateMember)
	r.GET("/update/member/:mail/:phoneNumber", CreateMember)
	r.POST("/update/member", CreateMember)
	r.GET("/read/member/:mail/:phoneNumber", ReadMember)
}

func CreateMember(c *gin.Context) {

	resData := NewJRes()

	m := model.Member{}
	err := c.ShouldBind(&m)
	if err != nil {
		fmt.Println(err)
		return
	}

	current := time.Now()
	m.CreateTime = current
	m.UpdateTime = current
	m.Enable = true

	_, err = coll.InsertOne(context.TODO(), m)
	if err != nil {
		resData["API-Call"] = "reject"
		resData["Reject-Reson"] = err
		c.JSON(http.StatusBadRequest, resData)
	}

	c.JSON(http.StatusOK, resData)
}

func UpdateMember(c *gin.Context) {

	m := model.Member{}
	err := c.ShouldBind(&m)
	if err != nil {
		fmt.Println(err)
		return
	}

	current := time.Now()
	m.CreateTime = current
	m.UpdateTime = current

	_, err = coll.InsertOne(context.TODO(), m)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ReadMember(c *gin.Context) {

	m := model.Member{}
	filter := model.MemberFilter{}

	err := c.ShouldBind(&filter)
	if err != nil {
		fmt.Println(err)
		return
	}

	cur, err := coll.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = cur.Decode(m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, m)
	}
}

func DeleteMember(c *gin.Context) {

	m := model.Member{}
	err := c.ShouldBind(&m)
	if err != nil {
		fmt.Println(err)
		return
	}

	current := time.Now()
	m.CreateTime = current
	m.UpdateTime = current

	_, err = coll.InsertOne(context.TODO(), m)
	if err != nil {
		fmt.Println(err)
		return
	}
}
