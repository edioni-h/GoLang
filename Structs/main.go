package main

import "fmt"

//this is how structs are defined

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

//Go interprets the properties from the type on the order given to it
//so it should go first name, goes on the first detail, and after the comma, the second property
func main() {
	// This is the first form: alex := person{"Alex", "Anderson"}
	//  alex := person{firstName: "Alex", lastName: "Anderson"}
	//  fmt.Print(alex)
	// var alex person

	// //this is how to give values to properties in a var with a type
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName:   "Jim",
		lastName:    "Troni",
		contactInfo: contactInfo{email: "jim@gmail.com", zipCode: 70000},
	}

	//#1 -so to create a pointer, first create a variable assigned the value := & var
	//jimPointer := &jim //turn value into address // *Type - turn address into value

	//#3 - and then to update the details, you also give the data in the said order

	//if the variable is already assigned the type data before, you can just use it
	//with the pointer function as a shortcut straight away

	jim.updateName("Jimmy", "Anderson")
	jim.print()

}

//#2 - then create a variable with *type of data, inside the new function with empty variables, and
//then to activate that type data into the variable, use (*variable).typeData =
// = the empty variables with an order
func (pointerToPerson *person) updateName(newFirstName string, newLastName string) {
	(*pointerToPerson).firstName = newFirstName
	(*pointerToPerson).lastName = newLastName
}

func (p person) print() {
	fmt.Printf("%v", p)
}
