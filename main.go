package main

import (
	"github.com/HtLord/ontnua/api"
	"github.com/HtLord/ontnua/db/mongo"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main()  {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Ontnua API")
	})

	api.MemberRouting(r)

	defer mongo.CloseDB()

	log.Fatal(r.Run()) // listen and serve on 0.0.0.0:8080
}
