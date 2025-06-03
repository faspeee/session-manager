package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"session-manager/configuration"
	"session-manager/controller"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	route := gin.Default()
	route.POST("/session/register", controller.Register)
	route.POST("/session/login", controller.Login)
	route.POST("/session/checkToken", controller.CheckToken)
	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := route.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/

	authorized.POST("admin", func(context *gin.Context) {
		user := context.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if context.Bind(&json) == nil {
			db[user] = json.Value
			context.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return route
}

func main() {
	configuration.CreateAndCloseTestContainerMongo()
	//configuration.Init()
	route := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	err := route.Run(":8282")
	if err != nil {
		return
	}
}
