// Value Types : int, string, bool, float, structs
// Reference Types : pointers, functions, map, channels, slices

package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// sau := person{"Saurabh", "Kale"}
	// a := person{firstName: "Alex", lastName: "Alexander"}
	// fmt.Println(a)
	// fmt.Println(sau)
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Alexander"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@email.com",
			zipCode: 94000,
		},
	}
	//jimPointer := &jim
	//jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
