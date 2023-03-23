package main

import (
	"fmt"
	"line-bot-app/linebot"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", linebot.Handler)
	fmt.Println("http://localhost:8080 で起動中...")
	// HTTPサーバを起動
	log.Fatal(http.ListenAndServe(":8080", nil))
	//database.Read("パスタ", database.Connect())
	//database.Migration()
	//apiClient.CreateRecipeMaster()
	//linebot.SendRecipeInfo("パスタ")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World!!!!"
	fmt.Fprintf(w, msg)
}
