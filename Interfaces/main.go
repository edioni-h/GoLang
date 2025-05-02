package main

import "fmt"

type englishBot struct {
}

type spanishBot struct {
}

//How to understand interfaces
/*So you create an interface that takes a function presented
in numerous scenarios/functions that take data from different types,
and if they fulfill the requirement, have the same function,
you can take it into an interface, that transforms all the types into
the specific type of the interface, which can later be used as that
type exactly in another function and call the different types
presented into the interface without rewriting code over n over
*/

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//this hidden code is useless
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (eb englishBot) getGreeting() string {
	//VERY custom logic for generating an english greeting

	return "Hi there!"

}

func (sb spanishBot) getGreeting() string {

	return "Hola!"
}
