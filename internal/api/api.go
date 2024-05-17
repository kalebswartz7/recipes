package api

import (
	"go/recipes/db"
	"go/recipes/models"
)

func AddRecipe(c *db.TRecipes, title string, description string, difficulty int, minutes int) (int, error) {
	recipe := models.Recipe{Title: title, Description: description, Difficulty: difficulty, Minutes: minutes}
	return c.InsertRecipe(recipe)
}

func AddStep(c *db.TRecipes, recipeId int, sequence int, category string, description string, minutes int) (int, error) {
	step := models.Step{RecipeId: recipeId, Sequence: sequence, Category: category, Description: description, Minutes: minutes}
	return c.InsertStep(step)
}

func AddIngredient(c *db.TRecipes, name string, category string) (int, error) {
	ingredient := models.Ingredient{Name: name, Category: category}
	return c.InsertIngredient(ingredient)
}

func AddQuantity(c *db.TRecipes, category string, amount float32) (int, error) {
	quantity := models.Quantity{Category: category, Amount: amount}
	return c.InsertQuantity(quantity)
}

func AddStepIngredient(c *db.TRecipes, stepId int, ingredientId int, quantityId int) (int, error) {
	stepIngredient := models.StepIngredient{StepId: stepId, IngredientId: ingredientId, QuantityId: quantityId}
	return c.InsertStepIngredient(stepIngredient)
}
