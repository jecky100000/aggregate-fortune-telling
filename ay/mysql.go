package ay

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	var yaml Yaml
	yaml.GetConf()
	dsn := yaml.Mysql.User + ":" + yaml.Mysql.Password + "@tcp(" + yaml.Mysql.Localhost + ":" + yaml.Mysql.Port + ")/" + yaml.Mysql.Database + "?parseTime=true&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败：", err)
	}
	Db = database

	//database.SingularTable(true)
	// 把模型与数据库中的表对应起来
	//Db.AutoMigrate(&models.Area{})
}
