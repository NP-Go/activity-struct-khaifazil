package main

import "fmt"

//Insert struct declaration here
type customer struct {
	fName string
	lName string
	age int
	subscriber bool
	address string
	phone int
	creditAvailable float32
	currentCartCost float32
	currentOrderCost float32
}

func main() {
	//Insert code here
	customer1 := customer {
		fName: "Annakin",
		lName: "Skywalker",
		age: 45,
		subscriber: true,
		address: "Death Star",
		phone: 1234567,
		creditAvailable: 10000.00,
		currentCartCost: 0.00,
		currentOrderCost: 0.00,
	}

	customer2 := customer {
		fName: "Han",
		lName: "Solo",
		age: 50,
		subscriber: false,
		address: "Tatooine",
		phone: 4321765,
		creditAvailable: 20000.00,
		currentCartCost: 0.00,
		currentOrderCost: 0.00,
	}
	
	fmt.Println(customer1, customer2)
}

