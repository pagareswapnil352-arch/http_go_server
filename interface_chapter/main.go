package main

import "fmt" // ← "fmt" = format package, gives us printing

// Interface (the rule)
type Animal interface {
	Sound() string
}

// Dog
type Dog struct{}

func (d Dog) Sound() string { return "Woof" }

// Cat
type Cat struct{}

func (c Cat) Sound() string { return "Meow" }

// Cow
type Cow struct{}

func (c Cow) Sound() string { return "Moo" }

// Function that prints any Animal's sound
func makeItSpeak(a Animal) {
	fmt.Println(a.Sound()) // 🖨️ PRINT happens here (inside the function)
}

func main() {
	dog := Dog{}
	cat := Cat{}
	cow := Cow{}

	fmt.Println("--- Animals speaking ---") // 🖨️ PRINT #1 (title)

	makeItSpeak(dog) // 🖨️ PRINT #2 → "Woof"
	makeItSpeak(cat) // 🖨️ PRINT #3 → "Meow"
	makeItSpeak(cow) // 🖨️ PRINT #4 → "Moo"
}
