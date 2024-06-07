package main

import (
	"bufio" // for files
	"fmt"
	"log" // for files
	"os" // for files
	 "strings"
	 "strconv" // for converting string to int
)

func main() {
	recipeList := getRecipeList()
	fmt.Println("Welcome to the Recipe Book!")
	choice := 10
	for choice != 0 {
		fmt.Print("\nHere are your options:\n1) View your recipies\n2) Add a recipe\n3) Edit a recipe\n4) See what you can bake\n0) Quit\nEnter your choice: ")
		fmt.Scanf("%d", &choice)
		switch choice {
			case 1: {
				for i := range recipeList {
					printRecipe(recipeList[i])
				}
			}
			case 2: {
				var newRecipe Recipe
				fmt.Print("\nEnter the name of the baked good: ")
				reader := bufio.NewReader(os.Stdin)
				n, err := reader.ReadString('\n')
				n = strings.Split(n, "\n")[0]
				if err != nil {
					fmt.Println(err)
				}
				newRecipe.name = n
				fmt.Println("Enter the ingredients (measurement unit ingredient):");
				ingredientCount := 1;
				for {
					var m float32
					var u string
					var i string
					fmt.Printf("%v) ", ingredientCount)
					fmt.Scanf("%g %s", &m, &u)
					reader := bufio.NewReader(os.Stdin)
					i, err := reader.ReadString('\n')
					i = strings.Split(i, "\n")[0] // removes the newline at the end of the string
					if err != nil {
						fmt.Println(err)
					}
					ingr := Ingredient{i, m, u}
					if ingr.name != "" {
						ingredientCount++;
							newRecipe.ingredientList = append(newRecipe.ingredientList, ingr)
					} else {
						break
					}
				}
				fmt.Print("Enter the baking temperature (F): ")
				var temp int
				fmt.Scanf("%d", &temp)
				newRecipe.bakingTempF = temp
				fmt.Print("Enter the baking time (minutes): ")
				var time int
				fmt.Scanf("%d", &time)
				newRecipe.bakingTimeMinutes = time
				recipeList = append(recipeList, newRecipe)
			}
			case 3: {
				fmt.Print("\nHere are your recipes:")
				for i, recipe := range recipeList {
					fmt.Printf("\n%v) %v", i + 1, recipe.name)
				}
				fmt.Print("\nEnter the recipe to edit: ")
				var recipeToEdit int
				fmt.Scanf("%d", &recipeToEdit)
				recipeList = editRecipeMenu(recipeList, recipeToEdit - 1)
			}
			case 4: {
				
			}
			default: {
				fmt.Println("\nInvalid input.")
			}
			case 0: {
				writeRecipesToFile(recipeList)
				fmt.Println("\nThank you for using my program.")
			}
		}
	}
}

func editIngredientMenu(list []Ingredient, index int) []Ingredient {
	choice := 10
	for choice != 0 {
		fmt.Print("\nHere are your options:\n1) Delete this ingredient\n2) Change name\n3) Change measurement\n0) Back\nEnter your choice: ")
		fmt.Scanf("%d", &choice)
		switch choice {
			case 1: {
			  list = append(list[0:index], list[index+1:]...)
				return list
			}
			case 2: {
				fmt.Print("\nEnter the new name: ")
				var newName string
				fmt.Scanf("%v", &newName)
				list[index].name = newName
			}
			case 3: {
				fmt.Print("\nEnter the new measurement: ")
				var newMeasure float32
				var newUnit string
				fmt.Scanf("%g %v", &newMeasure, &newUnit)
				list[index].measure = newMeasure
				list[index].measureUnit = newUnit
			}
			default:
				fmt.Println("\nInvalid input.")
			case 0:
		}
	}
	return list
}

func editRecipeMenu(list []Recipe, index int) []Recipe {
	choice := 10
	for choice != 0 {
		fmt.Print("\nHere are your options:\n1) Delete this recipe\n2) Edit an ingredient\n3) Edit baking time\n4) Edit baking temp\n0) Back\nEnter your choice: ")
		fmt.Scanf("%d", &choice)
		switch choice {
			case 1: {
				list = append(list[0:index], list[index+1:]...)
				// delete the recipe?
				return list
			}
			case 2: {
				fmt.Println("\nHere are your ingredients:")
				for i, ingredient := range list[index].ingredientList {
					fmt.Printf("%v) %s\n", i + 1, ingredient.name)
				}
				fmt.Print("Enter the ingredient to edit: ")
				var ingredientToEdit int
				fmt.Scanf("%d", &ingredientToEdit)
				editIngredientMenu(list[index].ingredientList, ingredientToEdit - 1)
			}
			case 3:
			case 4:
			default:
				fmt.Println("\nInvalid input.")
			case 0:
				return list
		}
	}
	return list
}

func getRecipeList() []Recipe { // read line by line https://www.geeksforgeeks.org/how-to-read-a-file-line-by-line-to-string-in-golang/
	var rList []Recipe
	file, err := os.Open("recipes.txt")
	if err != nil {
		log.Fatalf("failed to open: %v", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		if strings.Index(scanner.Text(), "Name: ") != -1 {
			var rec Recipe
			rec.name = strings.Split(scanner.Text(), "Name: ")[1]
			rList = append(rList, rec)
		} else if strings.Index(scanner.Text(), "Bake temp: ") != -1 {
			bTempString := strings.Split(scanner.Text(), "Bake temp: ")[1]
			bTemp, err := strconv.Atoi(bTempString)
			if err != nil {
				fmt.Println("failed to convert string to int")
			}
			rList[len(rList) - 1].bakingTempF = int(bTemp)
		} else if strings.Index(scanner.Text(), "Bake time: ") != -1 {
			bTimeString := strings.Split(scanner.Text(), "Bake time: ")[1]
			bTime, err := strconv.Atoi(bTimeString)
			if err != nil {
				fmt.Println("failed to convert string to int")
			}
			rList[len(rList) - 1].bakingTimeMinutes = int(bTime)
		} else if strings.Index(scanner.Text(), " - ") != -1 {
			ingredientString := strings.Split(scanner.Text(), " - ")[1]
			rList[len(rList) - 1].ingredientList = append(rList[len(rList) - 1].ingredientList, toIngredient(ingredientString))
		}
		text = nil // clears the array
		text = append(text, scanner.Text()) // appends the next line
	}
	file.Close()
	return rList
}

func recipeToString(r Recipe) string {
	s := "Name: " + r.name + "\nBake time: " + strconv.Itoa(r.bakingTimeMinutes) + "\nBake temp: " + strconv.Itoa(r.bakingTempF) + "\nInredients:\n"
	for i := range r.ingredientList {
		s = s + " - " + ingredientToString(r.ingredientList[i])
	}
	return s
}

func ingredientToString(i Ingredient) string {
	s := strconv.FormatFloat(float64(i.measure), 'f', -1, 32) + " " + i.measureUnit + " " + i.name + "\n"
	return s
}

func printRecipe(r Recipe) {
	fmt.Printf("\n%v\nBake time: %v minutes\nBake temp: %v F\nInredients:\n", r.name, r.bakingTimeMinutes, r.bakingTempF)
		for i := range r.ingredientList {
		fmt.Print(" - ")
		printIngredient(r.ingredientList[i])
	}
}

func printIngredient(i Ingredient) {
	fmt.Println(i.measure, i.measureUnit, i.name)
}

func toIngredient(iString string) Ingredient {
	var i Ingredient
	list := strings.Split(iString, " ")
	m, err := strconv.ParseFloat(list[0], 32)
	if err != nil {
		fmt.Println("failed to convert string to int")
	}
	i.measure = float32(m)
	i.measureUnit = list[1]
	i.name = list[2]
	for word := 3; word < len(list); word++ {
		i.name = i.name + " " + list[word]
	}
	return i
}

func writeRecipesToFile(rList []Recipe) {
	file, err1 := os.Create("recipes.txt")
	if err1 != nil {
		fmt.Println(err1)
	}
	for i := 0; i < len(rList); i++ {
		l, err2 := file.WriteString(recipeToString(rList[i]))
		if err2 != nil {
			fmt.Println(l)
			fmt.Println(err2)
		}
	}
	err3 := file.Close()
	if err3 != nil {
		fmt.Println(err3)
	}
}

type Ingredient struct {
	name string
	measure float32
	measureUnit string
}

type Recipe struct {
	name string
	ingredientList []Ingredient
	bakingTempF int
	bakingTimeMinutes int
}