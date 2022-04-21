package main

import (
	"gin/ay"
	"gin/routers"
	"gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	r *gin.Engine
)

func main() {

	ay.Yaml = ay.InitConfig()
	ay.Sql()
	go ay.WatchConf()

	r = gin.Default()

	gin.SetMode(gin.DebugMode)

	//r.Use(Cors())
	//r.Use(Header())

	r = service.Set(r)

	//r.LoadHTMLGlob("views/**/**/*")
	r.StaticFS("/static/", http.Dir("./static"))

	r = routers.GinRouter(r)

	err := r.Run(":8080")
	if err != nil {
		panic(err.Error())
	}
}
