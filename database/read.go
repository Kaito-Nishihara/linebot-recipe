package database

import (
	"context"
	"line-bot-app/ent"
	"line-bot-app/ent/category"
	"log"
)

func Read(text string, client *ent.Client) string {

	ctx := context.Background()
	recipe, err := client.Category.Query().Where(category.CategoryName(text)).All(ctx)
	if err != nil {
		log.Fatalf("invalid CategoryName: %v", err)
	}
	log.Print(recipe)
	return recipe[0].CategoryId
}
