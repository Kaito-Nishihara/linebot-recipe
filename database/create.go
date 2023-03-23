package database

import (
	"context"
	"line-bot-app/ent"
	"line-bot-app/model"
	"log"
)

func Create(model model.Category, client *ent.Client) {

	ctx := context.Background()
	usr, err := client.Debug().Category.
		Create().
		SetCategoryId(model.CategoryId).
		SetCategoryName(model.CategoryName).
		SetCategoryUrl(model.CategoryUrl).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed create user: %v", err)
	}
	log.Printf("user: %+v", usr)

	log.Print("ent sample done.")
}
