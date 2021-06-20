package main

import (
	"myapp/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	// fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test2"
	// u.Email = "test2@gmail.com"
	// u.PassWord = "testtest2"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "test3"
	// u.Email = "test3@gmail.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u, _ := models.GetUser(1)
	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	user, _ := models.GetUser(2)
	user.CreateTodo("First Todo")
}
