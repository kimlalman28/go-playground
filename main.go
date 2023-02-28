// Package == Project == Workspace
// Executable - generates a file that we can run
// Reusable - Code used as 'helpers'
// package main -> go build -> main.exe (if we ran this file, the function named main would automatically be ran)
// package main: defines a package that can be complied and then executed - must have a func called main!
// package calculator: defines a package that can be used as a dependency (helper code/ reusable package)
package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	// cards := newDeck()
	// cards.shuffle()
	// cards.print()

	// jim := person{
	// 	firstName: "Jim",
	// 	lastName: "Jones",
	// 	contactInfo: contactInfo{
	// 		email: "jim@test.com",
	// 		zipCode: 11234,
	// 	},
	// }

	// jim.print()
	// jim.updateName("Jimmy")
	// jim.print()

	// GO will automatically turn variable into pointer if receiver is expecting a pointer
	// jim.updateNameWithPointer("Jimmy")
	// jim.print()

	colors := map[string]string{
		"red": "#ff0000",
		"green": "ff9999",
	}

	colors2 := make(map[int]string)
	colors2[10] = "#ffffff"

	fmt.Println(colors)
	fmt.Println(colors2)

	delete(colors, "red")
	fmt.Println(colors)

	animals := map[string]string{
		"dog": "woof",
		"cat": "meow",
		"pig": "oink",
	}

	printMap(animals)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) updateNameWithPointer(newFirstName string) {
	(*p).firstName = newFirstName
}

func printMap(m map[string]string) {
	for animal, sound := range m {
		fmt.Println("Animal:", animal, "Sound:", sound)
	}
}
