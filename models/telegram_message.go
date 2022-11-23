package models

import (
	"encoding/json"
	"github.com/MMMahmoodian/alarm/conf/telegram"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type TelegramMessage struct {
	Body      string `json:"body"`
	ChatId    string `json:"chat_id"`
	ParseMode string `json:"parse_mode"`
}

func (tm TelegramMessage) Handle() error {
	telegramApi := telegram.GetSendMessageUrl()
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id":    {tm.ChatId},
			"text":       {tm.Body},
			"parse_mode": {tm.ParseMode},
		})
	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return err
	}
	defer response.Body.Close()
	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram Response: %s", bodyString)

	return nil
}

func (tm TelegramMessage) ToBytes() ([]byte, error) {
	return json.Marshal(tm)
}

func (tm TelegramMessage) FromBytes(bytes []byte) (Message, error) {
	message := TelegramMessage{}
	if err := json.Unmarshal(bytes, &message); err != nil {
		return nil, err
	}
	return message, nil
}
