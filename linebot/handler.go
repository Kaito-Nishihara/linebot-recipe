package linebot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"line-bot-app/database"
	"line-bot-app/model"
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// BOTを初期化
	bot, err := linebot.New(
		"fbea75a2e1de0b4a7a18c1d38831c9d6",
		"H7Vm/YxTj87tSWre3Dpoo7ZHmPUcC8Qqw0e9/j4MfGImT1n+abn/Doz+riFdTMystEJjDRFWCoJRFukE+d3XycG0Mn4IvEo4zFwQDr4dsLzN777ENqn+Cakr7wlzrO8W27gMN1XGwoe+V2MHT0FRhQdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		log.Fatal(err)
	}

	// リクエストからBOTのイベントを取得
	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		// イベントがメッセージの受信だった場合
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			// メッセージがテキスト形式の場合
			case *linebot.TextMessage:
				replyMessage := message.Text
				replyMsg := sendRecipeInfo(replyMessage)

				res := linebot.NewTemplateMessage(
					"レシピ一覧",
					linebot.NewCarouselTemplate(replyMsg...).WithImageOptions("rectangle", "cover"),
				)
				if _, err := bot.ReplyMessage(event.ReplyToken, res).Do(); err != nil {
					log.Print(err)
				}

			}
		}
	}
}

func sendRecipeInfo(text string) []*linebot.CarouselColumn {
	categoryId := database.Read(text, database.Connect())

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest("GET", "https://app.rakuten.co.jp/services/api/Recipe/CategoryRanking/20170426?format=json&categoryId="+categoryId+"&applicationId=1097504700709523230", nil)
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
	// fmt.Println(resp.Status)
	// fmt.Println(string(body))
	data := model.RecipeResult{}
	err = json.Unmarshal(body, &data)
	fmt.Println(data)
	var ccs []*linebot.CarouselColumn
	for _, item := range data.Result {

		cc := linebot.NewCarouselColumn(
			item.SmallImageUrl,
			item.RecipeTitle,
			item.RecipeCost,
			linebot.NewURIAction("楽天レシピで開く", item.RecipeUrl),
		).WithImageOptions("#FFFFFF")
		ccs = append(ccs, cc)
	}
	return ccs
}

// response APIレスポンス
