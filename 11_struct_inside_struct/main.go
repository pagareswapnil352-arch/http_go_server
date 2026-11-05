package main

import "fmt"

// 1. Contact struct
type Contact struct {
	Phone string
	Email string
}

// 2. Detail struct
type Detail struct {
	Name string
	Age  int
}

// 3. Address struct
type Address struct {
	City  string
	State string
}

// 4. Merge all three into Employee
type Employee struct {
	Detail  Detail  // struct inside struct
	Contact Contact // struct inside struct
	Address Address // struct inside struct
}

func main() {
	// Create the Employee by nesting each struct
	emp := Employee{
		Detail: Detail{
			Name: "Rahul",
			Age:  25,
		},
		Contact: Contact{
			Phone: "9876543210",
			Email: "rahul@example.com",
		},
		Address: Address{
			City:  "Mumbai",
			State: "Maharashtra",
		},
	}

	// Access nested fields with dots
	fmt.Println("Name  :", emp.Detail.Name)
	fmt.Println("Age   :", emp.Detail.Age)
	fmt.Println("Phone :", emp.Contact.Phone)
	fmt.Println("Email :", emp.Contact.Email)
	fmt.Println("City  :", emp.Address.City)
	fmt.Println("State :", emp.Address.State)

	// Modify a nested field
	emp.Address.City = "Pune"
	fmt.Println("New City:", emp.Address.City)
}
