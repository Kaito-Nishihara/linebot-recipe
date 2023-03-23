package database

import (
	"fmt"
	"line-bot-app/ent"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *ent.Client {
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		"localhost", "5432", "nishihara", "linebot", "masaki46C"))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	//defer client.Close()
	return client
}
