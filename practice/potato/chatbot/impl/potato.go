package impl

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jou66jou/go-myresume/practice/potato/chatbot"
)

//PotatoChatBot potao instance
type PotatoChatBot struct {
	apiURL string
	token  string
	client *http.Client
}

//GetID return id
func (c *PotatoChatBot) GetID() string {
	return "potato"
}

// NewPotoBotClient 提供新增 potato chatbot 客戶端
func NewPotoBotClient(apiURL string, botToken string, meta map[string]string) chatbot.PotatoClient {
	client := PotatoChatBot{
		apiURL: fmt.Sprintf("%s/%s", apiURL, botToken),
		token:  botToken,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Timeout: 5 * time.Second,
		},
	}
	return &client
}

func (c *PotatoChatBot) getGroups() (g *chatbot.Groups, err error) {
	req, _ := http.NewRequest("GET", c.apiURL+"/getGroups", nil)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(resp.Status)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	items := new(chatbot.Groups)
	err = json.Unmarshal(b, items)
	if err != nil {
		return nil, err
	}
	return items, nil

}

// SetWebhook 提供設定 potato chatbot webhook
func (c *PotatoChatBot) SetWebhook(opts chatbot.PotatoWebhookOpts) (err error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, c.apiURL+"/setWebhook", bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		return fmt.Errorf(resp.Status)
	}
	return nil
}

// RespUpdateMessage 提供回覆 potato chatbot 的訊息
func (c *PotatoChatBot) RespUpdateMessage(opts chatbot.SendTextMessage) (err error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, c.apiURL+"/sendTextMessage", bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		var respBody map[string]interface{}
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &respBody)
		if err != nil {
			return err
		}
		return fmt.Errorf(resp.Status + " : " + string(b))
	}
	return nil
}

// BroadcastToGroup 提供 potato chatbot 廣播訊息到群組的實作
func (c *PotatoChatBot) BroadcastToGroup(msg string) (err error) {
	g, err := c.getGroups()
	fmt.Printf("%+v\n", g)
	if err != nil {
		err = fmt.Errorf("get groups err: %s", err.Error())
		return err
	}
	if g == nil || (g.Items.Channels == nil && g.Items.Groups == nil) {
		err = errors.New("channels not found in this group")
		return err
	}
	url := c.apiURL + "/sendTextMessage"

	var req *http.Request

	for _, i := range g.Items.Channels {
		sendMsg := chatbot.SendTextMessage{
			ChatType: chatbot.PotatoChartTypeChannel,
			ChatID:   i.ID,
			Text:     msg,
		}
		body, err := json.Marshal(sendMsg)
		if err != nil {
			return err
		}
		req, err = http.NewRequest("POST", url, bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		if err == nil {
			err = c.doRequest(req)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
	for _, i := range g.Items.Groups {
		sendMsg := chatbot.SendTextMessage{
			ChatType: chatbot.PotatoChartTypeGroup,
			ChatID:   i.ID,
			Text:     msg,
		}
		body, err := json.Marshal(sendMsg)
		if err != nil {
			return err
		}
		req, err = http.NewRequest("POST", url, bytes.NewReader(body))
		req.Header.Add("Content-Type", "application/json")
		if err == nil {
			err = c.doRequest(req)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}

	return nil
}

func (c *PotatoChatBot) doRequest(req *http.Request) (err error) {
	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return fmt.Errorf("status:%s -- %s", resp.Status, string(b))
	}

	return nil
}
