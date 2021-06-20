package main

import (
	"fmt"
	"myapp/app/models"
)

func main() {
	t, _ := models.GetTodo(1)
	fmt.Println(t)
}
