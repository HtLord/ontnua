package main

import (
	"fmt"
	"github.com/HtLord/ontnua/api"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main()  {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Hello", "<h1>Hello world MTFK!</h1>")
	})
	r.POST("/create/keeper", api.CreateTest)
	fmt.Println("Start serving")
	defer api.CloseDB()
	log.Fatal(r.Run()) // listen and serve on 0.0.0.0:8080
}
