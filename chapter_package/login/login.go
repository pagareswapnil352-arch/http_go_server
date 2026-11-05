package login

import "fmt"

type Details struct {
	Name string

	Username string

	Age int
}

func (d Details) Login() {

	fmt.Println("the username :", d.Name)
	fmt.Println("the username :", d.Age)
	fmt.Println("the username :", d.Username)

}

func Greeting(name string) {

	fmt.Println("the greetings are:", name)
}
