package main

import (
	"fmt"

	"github.com/ogarciamartinez/interfacesample/animal"
)

// Dog animal
type Dog struct {
}

// Speak refers to Dog animal
func (d Dog) Speak() string {
	return "Woof!"
}

// Cat animal
type Cat struct {
}

// Speak refers to Cat animal
func (c Cat) Speak() string {
	return "Meow!"
}

// JavaProgrammer animal
type JavaProgrammer struct {
}

// Speak refers to JavaProgrammer animal
func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func main() {
	animals := []animal.Animal{Dog{}, Cat{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
