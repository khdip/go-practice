package main

import (
	"fmt"
	"strconv"
)

type employee struct {
	firstName  string
	lastName   string
	department string
	position   string
	salary     int
}

func main() {
	employee1 := employee{firstName: "John", lastName: "Doe", department: "Software", position: "php Developer", salary: 45000}
	employee2 := employee{"Amanda", "Doe", "Digital", "Graphic Designer", 30000}

	fmt.Println(employee1)
	fmt.Println(employee2)
	fmt.Println("Salary of " + employee2.firstName + " is " + strconv.Itoa(employee2.salary) + ".")
}
