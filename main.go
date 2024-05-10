package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Recipe Book!")
	choice := 10;
	for choice != 0 {
		fmt.Print("Here are your options:\n1) View your recipies\n2) Add a recipe\n3) Edit a recipe\n4) See what you can bake\n0) Quit\nEnter your choice: ")
		fmt.Scanf("%d", &choice)
		switch choice {
			case 1:
			case 2:
			fmt.Println("Enter the ingredients (measurement unit ingredient):");
			ingredientCount := 1;
			for {
				var m float32
				var u string
				var i string
				fmt.Printf("%v) ", ingredientCount)
				fmt.Scanf("%g %s %s", &m, &u, &i)
				ingr := Ingredient{i, m, u}
				fmt.Println(ingr)
				ingredientCount++;
				if ingr.name == "" {
					break
				}
			}
			case 3:
			case 4:
			case 0:
		}
	}
}

type Ingredient struct {
	name string
	measure float32
	measureUnit string
}

type Recipe struct {
	ingredientList []Ingredient
	bakingTemp int
	bakingTime int
}