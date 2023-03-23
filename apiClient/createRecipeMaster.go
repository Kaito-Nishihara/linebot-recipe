package apiClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"line-bot-app/database"
	"line-bot-app/model"
	"net/http"
)

func CreateRecipeMaster() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest("GET", "https://app.rakuten.co.jp/services/api/Recipe/CategoryList/20170426?format=json&applicationId=1097504700709523230", nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
	fmt.Println(string(body))
	data := model.Result{}
	err = json.Unmarshal(body, &data)
	clientDatebase := database.Connect()
	fmt.Println(data)
	for i, v := range data.Result.Large {
		database.Create(v, clientDatebase)
		fmt.Println(i)
	}

	for i, v := range data.Result.Medium {
		database.Create(v, clientDatebase)
		fmt.Println(i)
	}

	for i, v := range data.Result.Small {
		database.Create(v, clientDatebase)
		fmt.Println(i)
	}
}
