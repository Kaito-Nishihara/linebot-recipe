package model

type RecipeResult struct {
	Result []Recipe `json:"result"`
}

type Recipe struct {
	FoodImageUrl      string   `json:"foodImageUrl"`
	MediumImageUrl    string   `json:"mediumImageUrl"`
	Nickname          string   `json:"nickname"`
	Pickup            string   `json:"pickup"`
	Rank              string   `json:"rank"`
	RecipeCost        string   `json:"recipeCost"`
	RecipeDescription string   `json:"recipeDescription"`
	RecipeId          int64    `json:"recipeId"`
	RecipeIndication  string   `json:"recipeIndication"`
	RecipeMaterial    []string `json:"recipeMaterial"`
	RecipePublishday  string   `json:"recipePublishday"`
	RecipeTitle       string   `json:"recipeTitle"`
	RecipeUrl         string   `json:"recipeUrl"`
	Shop              int64    `json:"shop"`
	SmallImageUrl     string   `json:"smallImageUrl"`
}
