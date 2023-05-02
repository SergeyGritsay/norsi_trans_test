package mainimport

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"posts_api/controllers"
	"posts_api/routes"

	"posts_api/initializers"
)

var (
	server *gin.Engine

	PostController      controllers.PostController
	PostRouteController routes.PostRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load env", err)
	}

	initializers.ConnectDB(&config)
	PostController = controllers.NewPostController(initializers.DB)
	PostRouteController = routes.NewRoutePostController(PostController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load env", err)
	}

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome"
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": message,
		})
	})

	PostRouteController.PostRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
