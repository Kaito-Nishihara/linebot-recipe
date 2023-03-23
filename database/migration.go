package database

import (
	"context"
	"fmt"
	"line-bot-app/ent"
	"line-bot-app/ent/migrate"
	"log"
)

func Migration() {
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		"localhost", "5432", "nishihara", "linebot", "masaki46C"))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	if err := client.Schema.Create(ctx, migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
	log.Print("ent sample done.")
}
