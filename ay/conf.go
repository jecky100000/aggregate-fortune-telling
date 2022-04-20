package ay

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var Yaml *viper.Viper

func initConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath("conf/")
	config.SetConfigType("yaml")
	err := config.ReadInConfig()
	if err != nil {
		log.Printf("Failed to get the configuration.")
	}
	return config
}

func watchConf() {
	Yaml.WatchConfig()
	Yaml.OnConfigChange(func(event fsnotify.Event) {
		// 配置文件修改重新执行的方法
		sql()
		//
		log.Printf("Detect config change: %s \n", event.String())
	})
}
