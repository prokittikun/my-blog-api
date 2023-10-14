package main

import (
	"fmt"
	config "my-blog-api/config"
	models "my-blog-api/models"
	routes "my-blog-api/routes"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
		return
	}
	fmt.Println("Connect SUCCESS")
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Blog{})
	r := routes.SetupRouter()
	//running
	r.Run()
}
