package main

import (
	"aggregate-fortune-telling/ay"
	"aggregate-fortune-telling/routers"
	"aggregate-fortune-telling/service"
	"aggregate-fortune-telling/task"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"net/http"
)

var (
	r *gin.Engine
)

func main() {

	ay.Yaml = ay.InitConfig()
	ay.Sql()
	go ay.WatchConf()

	// 开启定时任务
	c := cron.New()

	c.AddFunc("@every 10m", task.Start)
	c.Start()

	gin.SetMode(gin.ReleaseMode)

	r = gin.Default()
	r = service.Set(r)
	r.StaticFS("/static/", http.Dir("./static"))
	r = routers.GinRouter(r)

	err := r.Run(":" + ay.Yaml.GetString("server.port"))
	if err != nil {
		panic(err.Error())
	}
	select {}
}
