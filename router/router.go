package router

import (
	"github.com/bonggar/gorestapi/config"
	"github.com/bonggar/gorestapi/service"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

//Cors : handle client origin rules
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

//Make : create endpoints
func Make() *gin.Engine {
	//Default : debug
	gin.SetMode(config.GinMode)

	//Default : with the Logger and Recovery middleware already attached
	r := gin.Default()

	//Serve frontend static files, ex: ReactJS
	r.Use(static.Serve("/", static.LocalFile("./view", true)))

	//Applying CORS rule
	r.Use(Cors())

	//Grouping to Version 1.0
	v1 := r.Group("api/v1")
	{
		v1.GET("/users", service.GetUsers)
		v1.GET("/users/:id", service.GetUser)
		v1.POST("/users", service.CreateUser)
		v1.PUT("/users/:id", service.UpdateUser)
		v1.DELETE("/users/:id", service.DeleteUser)
		v1.OPTIONS("/users", service.OptionsUser)
		v1.OPTIONS("/users/:id", service.OptionsUser)
	}

	return r
}
