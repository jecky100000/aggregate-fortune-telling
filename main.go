package main

import (
	"gin/controllers/api"
	"gin/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var (
	r *gin.Engine
)

func main() {

	r = gin.Default()

	// gin.SetMode(gin.DebugMode)

	r.Use(Cors())
	r.Use(Header())

	//r = service.Set(r)

	//r.LoadHTMLGlob("views/**/**/*")
	r.StaticFS("/static/", http.Dir("./static"))

	r = routers.GinRouter(r)

	err := r.Run(":8080")
	if err != nil {
		panic(err.Error())
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			c.Header("Access-Control-Allow-Headers", "*")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}

func Header() gin.HandlerFunc {
	return func(c *gin.Context) {
		for name, values := range c.Request.Header {
			// Loop over all values for the name.
			for _, value := range values {
				if name == "Authorization" {
					api.Token = value
				}
				if name == "From" {
					api.Appid, _ = strconv.Atoi(value)
				}
			}
		}

		c.Next()
	}
}
