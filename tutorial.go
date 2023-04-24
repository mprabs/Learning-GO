package main

import "fmt"


func main() {
	todos := []string{}
	todo:=""
	add := true
	userControl:=""

	for add {
		
		fmt.Println("Add todo")
		fmt.Scan(&todo)
		todos = append(todos, todo)


		fmt.Println("Add another? Y or N ")
		fmt.Scan(&userControl)

		if userControl == "N" {
			add = false
		}
	}


	for item := range todos {
		fmt.Println(todos[item])
	}	
}

