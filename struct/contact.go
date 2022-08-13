package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type contactPerson struct {
	contactInfo
	firstName string
	lastName  string
}

func main() {
	jim := contactPerson{
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipcode: 239003,
		},
		firstName: "Jim",
		lastName:  "Anderson",
	}

	fmt.Println(jim)
}
