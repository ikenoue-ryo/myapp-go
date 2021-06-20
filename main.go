package main

import (
	"fmt"
	"myapp/app/models"
)

func main() {
	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// user, _ := models.GetUser(2)
	// user.CreateTodo("Second Todo")

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	user2, _ := models.GetUser(2)
	todos, _ := user2.GetTodosByUser()
	for _, v := range todos {
		fmt.Println(v)
	}
}
