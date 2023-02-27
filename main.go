// Package == Project == Workspace
// Executable - generates a file that we can run
// Reusable - Code used as 'helpers'
// package main -> go build -> main.exe (if we ran this file, the function named main would automatically be ran)
// package main: defines a package that can be complied and then executed - must have a func called main!
// package calculator: defines a package that can be used as a dependency (helper code/ reusable package)
package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
