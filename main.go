package main

import (
	"fmt"
	"myapp/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test2"
	u.Email = "test2@gmail.com"
	u.PassWord = "testtest2"
	fmt.Println(u)

	u.CreateUser()
}
