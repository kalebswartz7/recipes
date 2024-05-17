package db

import (
	"database/sql"
	"go/recipes/models"

	_ "github.com/mattn/go-sqlite3"
)

type TRecipes struct {
	db *sql.DB
}

const dbName string = "db/recipes.db"

func GetDb() (*TRecipes, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}
	return &TRecipes{
		db: db,
	}, nil
}

func (c *TRecipes) InsertRecipe(recipe models.Recipe) (int, error) {
	res, err := c.db.Exec("INSERT INTO recipes(title, description, difficulty, minutes) VALUES(?,?,?,?);",
		recipe.Title, recipe.Description, recipe.Difficulty, recipe.Minutes)
	if err != nil {
		return 0, err
	}

	return getInsertResult(res)
}

func (c *TRecipes) InsertStep(step models.Step) (int, error) {
	res, err := c.db.Exec("INSERT INTO steps(recipe_id, sequence, category, description, minutes) VALUES(?,?,?,?,?);",
		step.RecipeId, step.Sequence, step.Category, step.Description, step.Minutes)
	if err != nil {
		return 0, err
	}

	return getInsertResult(res)
}

func (c *TRecipes) InsertIngredient(ingredient models.Ingredient) (int, error) {
	res, err := c.db.Exec("INSERT INTO ingredients(name, category) VALUES(?,?);",
		ingredient.Name, ingredient.Category)
	if err != nil {
		return 0, err
	}

	return getInsertResult(res)
}

func (c *TRecipes) InsertQuantity(quantity models.Quantity) (int, error) {
	res, err := c.db.Exec("INSERT INTO quantities(category, amount) VALUES(?,?);",
		quantity.Category, quantity.Amount)
	if err != nil {
		return 0, err
	}

	return getInsertResult(res)
}

func (c *TRecipes) InsertStepIngredient(stepIngredient models.StepIngredient) (int, error) {
	res, err := c.db.Exec("INSERT INTO step_ingredients(step_id, ingredient_id, quantity_id) VALUES(?,?,?);",
		stepIngredient.StepId, stepIngredient.IngredientId, stepIngredient.QuantityId)
	if err != nil {
		return 0, err
	}

	return getInsertResult(res)
}

func getInsertResult(res sql.Result) (int, error) {
	var id int64
	var err error

	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}
