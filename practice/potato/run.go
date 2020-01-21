package main

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/jou66jou/go-myresume/practice/potato/chatbot"
	"github.com/jou66jou/go-myresume/practice/potato/chatbot/impl"
)

var (
	apiURL = "https://api.rct2008.com:8443"
	token  = "10260663:9RNP75o8J1v6TDHZj65eqj8g"
)

type User struct {
	Name string
}

func main() {
	// Potato 對帳用聊天機器人 Client
	potatoPublicClient := impl.NewPotoBotClient(apiURL, token, nil)
	// 設定 webhook 接收 potato 訊息
	webhookOpts := chatbot.PotatoWebhookOpts{
		URL: "https://ptsv2.com/t/testPotato/post",
		// Certificate: "",
	}
	err := potatoPublicClient.SetWebhook(webhookOpts)
	if err != nil {
		fmt.Println("SetWebhook : ", err)
		return
	}

	// 使用 teamplate
	u := User{"Group"}
	tmpl, err := template.New("test").Parse(chatbot.TestMsg)
	if err != nil {
		fmt.Println("template.New : ", err)
		return
	}
	var resp bytes.Buffer
	err = tmpl.Execute(&resp, u)
	if err != nil {
		fmt.Println("tmpl.Execute : ", err)
		return
	}

	// // 傳送個人訊息
	// sendMsgOpts := chatbot.SendTextMessage{
	// 	ChatType: 1,
	// 	ChatID:   22981873,
	// 	Text:     resp.String(),
	// 	// ReplyToMessageID: 3,
	// }
	// err = potatoPublicClient.RespUpdateMessage(sendMsgOpts)
	// if err != nil {
	// 	fmt.Println("RespUpdateMessage : ", err)
	// 	return
	// }

	// 傳送群組訊息
	err = potatoPublicClient.BroadcastToGroup(resp.String())
	if err != nil {
		fmt.Println("BroadcastToGroup : ", err)
		return
	}
	fmt.Println("done.")
}
