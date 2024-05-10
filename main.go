package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Recipe Book!")
	choice := 10;
	for choice != 0 {
		fmt.Print("Here are your options:\n1) View your recipies\n2) Add a recipe\n3) Edit a recipe\n4) See what you can bake\n0) Save and quit\nEnter your choice: ")
		fmt.Scanf("%d", &choice)
		switch choice {
			case 1:
			case 2:
			case 3:
			case 4:
			case 0:
		}
	}
}

type Ingredient struct {
	ingredientName string
	measure float32
	measureUnit string
}

type Recipe struct {
	ingredientList []Ingredient
	bakingTemp int
	bakingTime int
}