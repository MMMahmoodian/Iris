package telegram

import "os"

func GetSendMessageUrl() string {
	return getBaseUrl() + "/sendMessage"
}

func getBaseUrl() string {
	return os.Getenv("TELEGRAM_BOT_API_URL") + os.Getenv("TELEGRAM_BOT_API_KEY")
}
