package main

import "fmt"

const g_cap int = 5 //total capacity of our list
var g_groceries [g_cap]string
var g_len int = 0 //total no. of grocery items in our current list

func add_grocery(a string) {
	if g_len < g_cap {
		g_groceries[g_len] = a
		g_len++
	} else {
		fmt.Println("too many items,now we are done for!!")
	}
}

func list_groceries() {
	fmt.Println("grocery list is as follows:")
	for i := 0; i < g_len; i++ {
		fmt.Println(g_groceries[i])
	}
}

func main() {
	add_grocery("bread")
	add_grocery("cucumber")
	add_grocery("potato chips")
	add_grocery("ice cream")
	add_grocery("fruit cake")
	add_grocery("holiday")
	list_groceries()
}
