package main

import "fmt"

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func (prod *Product) printData() {
	fmt.Printf("ID: %v\n", prod.id)
	fmt.Printf("Title: %v\n", prod.title)
	fmt.Printf("Description: %v\n", prod.description)
	fmt.Printf("Price: USD %.2f\n", prod.price)
}

func NewProduct(id string, t string, d string, p float64) *Product {
	return &Product{id, t, d, p}
}

func main() {
	firstProduct := Product{
		"first-product",
		"A Book",
		"A nice book",
		10.99,
	}
	secondProduct := NewProduct(
		"second-product",
		"A Carpet",
		"A beautiful carpet",
		99.99,
	)

	firstProduct.printData()
	secondProduct.printData()
}

// Your Tasks
// 1) Create a new type / data structure for storing product data (e.g. about a book)
//		The data structure should contain these fields:
//		- ID
//		- Title / Name
//		- Short description
//		- price (number without currency, we'll just assume USD)
// 2) Create concrete instances of this data type in two different ways:
//		- Directly inside of the main function
//		- Inside of main, by using a "creation helper function"
//		(use any values for title etc. of your choice)
//		Output (print) the created data structures in the command line (in the main function)
// 3) Add a "connected function" that outputs the data + call that function from inside "main"
// 4) Change the program to fetch user input values for the different data fields
//		and create only one concrete instance with that entered data.
//		Output that instance data (via the connected function) then.
// 5) Bonus: Add a connected "store" function that writes that data into a file
//		The filename should be the unique id, the function should be called at the
//		end of main.
