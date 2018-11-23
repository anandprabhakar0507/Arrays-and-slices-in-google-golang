package main

import "fmt"

var g_groceries []string

func add_grocery(a string) {
	fmt.Println("capacity", cap(g_groceries))
	g_groceries = append(g_groceries, a)
}

func list_groceries() {
	fmt.Println("grocery list is as follows:")
	for i := 0; i < len(g_groceries); i++ {
		fmt.Println(g_groceries[i])
	}
}

func main() {
	add_grocery("bread")
	add_grocery("cucumber")
	add_grocery("potato chips")
	add_grocery("ice cream")
	add_grocery("fruit cake")
	list_groceries()
}
