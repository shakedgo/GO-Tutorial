package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact contactInfo
}

func main() {
	person1 := person{
		firstName: "Shaked",
		lastName: "Govri",
		contact: contactInfo{
			email: "shakedgo@gmail.com",
			zipCode: 12345,
		},
	}
	person1.lastName = "Govrin"
	// person1.changeName("Emanuel") // Doesn't work because it's just a pointer
	person1Pointer := &person1 //&=gets the memory location of the variable
	person1Pointer.changeName("Emanuel")
	person1.print()


}

// func (p person) changeName(newName string) {
// 	p.firstName = newName
// }

// *person = means we want the type to be a pointer to a person
func (pointerToPerson *person) changeName(newName string) {
	// *=give me access to the value in this memory location
	(*pointerToPerson).firstName = newName
}

func (p person) print()  {
	fmt.Println(p) 
	// // {Shaked Govrin {shakedgo@gmail.com 12345}}
	fmt.Printf("%+v", p)
	// // {firstName: Shaked lastName: Govrin, contact:{email: "shakedgo@gmail.com", zipCode: 12345}}
}