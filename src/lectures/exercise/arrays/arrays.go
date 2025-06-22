//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	Name  string
	Price float32
}

func main() {
	shoppingList := [4]Product{}
	shoppingList[0] = Product{Name: "Apple", Price: 3.33}
	shoppingList[1] = Product{Name: "Banana", Price: 2.99}
	shoppingList[2] = Product{Name: "Orange", Price: 3.33}
	fmt.Println("last item:\t", shoppingList[len(shoppingList)-1])
	fmt.Println("total items:\t", len(shoppingList))
	fmt.Println("total cost of items:\t", calculateTotal(shoppingList[:]))
	shoppingList[3] = Product{Name: "Tomatoes", Price: 4.44}
	fmt.Println("total items:\t", len(shoppingList))
	fmt.Println("total cost of items:\t", calculateTotal(shoppingList[:]))
}
func calculateTotal(p []Product) float32 {
	var out float32 = 0

	for _, item := range p {
		out += item.Price
	}

	return out
}
