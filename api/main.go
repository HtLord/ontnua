package back

import (
	"fmt"
	"github.com/HtLord/glacial-brook-86888/servapi"
	"github.com/HtLord/ontnua/back"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main()  {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Hello", "<h1>Hello world MTFK!</h1>")
	})
	r.POST("/create/keeper", back.GinCreateTest)
	fmt.Println("Start serving")
	defer servapi.CloseDB()
	log.Fatal(r.Run()) // listen and serve on 0.0.0.0:8080
}

func notGin(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>My API will work on Heroku. But not sure about db connection.</h1>")
	})
	http.HandleFunc("/create/keeper", servapi.CreateKeeper)
}