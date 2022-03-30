package main

import "fmt"

// ContactInfo struct
type ContactInfo struct {
	email   string
	zipCode int
}

// Note:(We can have a struct inside a struct)

// Person struct that is having ContactInfo struct in it. We can say it is a nested struct.
type Person struct {
	firstName   string
	lastName    string
	ContactInfo // nested struct
}

func main() {

	// declaring a variable `boy` of type Person struct and assigning value according to it's fields type
	var boy = Person{
		firstName: "Jimyy",
		lastName:  "bucky",
		ContactInfo: ContactInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}
	// you can declare the boy variable in this way aswell `boy := Person{}`

	// printing the struct
	fmt.Println("The struct values -> ", boy)

}

// Note:= If we use capital first letter of a struct then we are declaring them to be use as global.
