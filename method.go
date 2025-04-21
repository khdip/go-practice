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

// Method
func (emp employee) getSalary() string {
	return "Salary of " + emp.firstName + " is " + strconv.Itoa(emp.salary) + "."
}

func (emp *employee) salaryIncrement() {
	emp.salary += 5000
}

func (emp *employee) gotPromoted(newPosition string) {
	emp.position = newPosition
}

func main() {
	employee1 := employee{firstName: "John", lastName: "Doe", department: "Software", position: "php Developer", salary: 45000}
	employee2 := employee{"Amanda", "Doe", "Digital", "Graphic Designer", 30000}

	fmt.Println(employee1)
	fmt.Println(employee2)
	fmt.Println(employee1.salary)
	fmt.Println(employee2.getSalary())
	employee2.salaryIncrement()
	fmt.Println(employee2.getSalary())
	employee1.gotPromoted("Senior php Developer")
	fmt.Println(employee1.firstName + " got promoted to " + employee1.position + ".")
}
