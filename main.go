package main

import (
	"fmt"
)

func main() {
	var recipeList []Recipe
	fmt.Println("Welcome to the Recipe Book!")
	choice := 10;
	for choice != 0 {
		fmt.Print("\nHere are your options:\n1) View your recipies\n2) Add a recipe\n3) Edit a recipe\n4) See what you can bake\n0) Quit\nEnter your choice: ")
		fmt.Scanf("%d", &choice)
		switch choice {
			case 1: {
				for i := 0; i < len(recipeList); i++ {
					printRecipe(recipeList[i])
				}
			}
			case 2: {
				var newRecipe Recipe
				fmt.Print("Enter the name of the baked good: ")
				var n string
				fmt.Scanf("%s", &n)
				newRecipe.name = n
				fmt.Println("Enter the ingredients (measurement unit ingredient):");
				ingredientCount := 1;
				for {
					var m float32
					var u string
					var i string
					fmt.Printf("%v) ", ingredientCount)
					fmt.Scanf("%g %s %s", &m, &u, &i)
					ingr := Ingredient{i, m, u}
					if ingr.name != "" {
						ingredientCount++;
							newRecipe.ingredientList = append(newRecipe.ingredientList, ingr)
					} else {
						break
					}
				}
				fmt.Print("Enter the baking temperature: ")
				var temp int
				fmt.Scanf("%d", &temp)
					newRecipe.bakingTemp = temp
				fmt.Print("Enter the baking time: ")
				var time int
				fmt.Scanf("%d", &time)
				newRecipe.bakingTime = time
				recipeList = append(recipeList, newRecipe)
			}
			case 3:
			case 4:
			case 0: {
				fmt.Println("Thank you for using my program.")
			}
		}
	}
}

func printRecipe(r Recipe) {
	fmt.Printf("\n%v\nBake for %v minutes at %v F.\nInredients:\n", r.name, r.bakingTime, r.bakingTemp)
	for i := 0; i < len(r.ingredientList); i++ {
		fmt.Print(" - ")
		printIngredient(r.ingredientList[i])
	}
}

func printIngredient(i Ingredient) {
	fmt.Println(i.measure, i.measureUnit, i.name)
}

type Ingredient struct {
	name string
	measure float32
	measureUnit string
}

type Recipe struct {
	name string
	ingredientList []Ingredient
	bakingTemp int
	bakingTime int
}