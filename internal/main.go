package main

import (
	"fmt"
	"go/recipes/api"
	"go/recipes/db"
)

func main() {
	c, err := db.GetDb()
	if err != nil {
		fmt.Printf("Error opening DB - %v\n)", err)
	}

	tryAddStepIngredient(c)
}

func tryAddRecipe(c *db.TRecipes) {
	fmt.Println("Let's try to add a new recipe to our table...")
	id, err := api.AddRecipe(c, "My third recipe", "I hope this tastes good..", 10, 30)
	if err != nil {
		fmt.Printf("Was not able to add recipe :( - %v\n)", err)
		return
	}
	fmt.Printf("Successfully added recipe! id: %v\n", id)
}

func tryAddStep(c *db.TRecipes) {
	fmt.Println("Let's try to add a new step to our table...")
	id, err := api.AddStep(c, 1, 1, "Prep", "First step for first recipe", 30)
	if err != nil {
		fmt.Printf("Was not able to add recipe :( - %v\n)", err)
		return
	}
	fmt.Printf("Successfully added recipe! id: %v\n", id)
}

func tryAddIngredient(c *db.TRecipes) {
	fmt.Println("Let's try to add a new ingredient to our table...")
	id, err := api.AddIngredient(c, "lemon", "fruit")
	if err != nil {
		fmt.Printf("Was not able to add recipe :( - %v\n)", err)
		return
	}
	fmt.Printf("Successfully added recipe! id: %v\n", id)
}

func tryAddQuantity(c *db.TRecipes) {
	fmt.Println("Let's try to add a new quantity to our table...")
	id, err := api.AddQuantity(c, "cup", 1.5)
	if err != nil {
		fmt.Printf("Was not able to add quantity :( - %v\n)", err)
		return
	}
	fmt.Printf("Successfully added quantity! id: %v\n", id)
}

func tryAddStepIngredient(c *db.TRecipes) {
	fmt.Println("Let's try to add a new step_ingredient to our table...")
	id, err := api.AddStepIngredient(c, 1, 1, 1)
	if err != nil {
		fmt.Printf("Was not able to add step_ingredient :( - %v\n)", err)
		return
	}
	fmt.Printf("Successfully added step_ingredient! id: %v\n", id)
}
