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
	r.PUT("/create/member", CreateMember)
	r.POST("/update/member", UpdateMember)
	r.GET("/read/member/:name", ReadMember)
	r.DELETE("/delete/member/:mail/:phoneNumber/:secret", DeleteMember)
}

func CreateMember(c *gin.Context) {

	m := model.Member{}
	err := c.ShouldBind(&m)
	if err != nil {
		c.JSON(http.StatusBadRequest, RejectRes(err))
		return
	}

	current := time.Now()
	m.CreateTime = current
	m.UpdateTime = current
	m.Enable = true

	err = mongo.CreateOne(coll, &m)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, RejectRes(err))
		return
	}

	c.JSON(http.StatusOK, AcceptRes())
}

func UpdateMember(c *gin.Context) {

	filter := model.MemberFilter{}
	err := c.ShouldBind(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, RejectRes(err))
		return
	}

	m := model.Member{}
	err = c.ShouldBind(&m)
	if err != nil {
		c.JSON(http.StatusBadRequest, RejectRes(err))
		return
	}

	current := time.Now()
	m.UpdateTime = current

	_, err = coll.UpdateOne(context.TODO(), mongo.UpdateFilter(filter), m)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, RejectRes(err))
		return
	}

	c.JSON(http.StatusOK, AcceptRes())
}

func ReadMember(c *gin.Context) {

	m := model.Member{}
	filter := make(map[string]string)
	for _, v := range c.Params {
		filter[v.Key] = v.Value
	}

	fmt.Println(filter)

	err := mongo.ReadOne(coll, filter, &m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, RejectResWith("ReadMember using coll-operation ReadOne", err))
		return
	}

	c.JSON(http.StatusOK, m)
}

func DeleteMember(c *gin.Context) {

	filter := model.MemberDelete{}
	err := c.ShouldBind(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, RejectRes(err))
		return
	}

	m := model.Member{}
	err = c.ShouldBind(&m)
	if err != nil {
		c.JSON(http.StatusBadRequest, RejectRes(err))
		return
	}

	current := time.Now()
	m.UpdateTime = current
	m.Enable = false

	_, err = coll.UpdateOne(context.TODO(), mongo.UpdateFilter(filter), m)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, RejectRes(err))
		return
	}

	c.JSON(http.StatusOK, AcceptRes())
}
