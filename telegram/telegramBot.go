package telegram
import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type TelegramBot struct {
	bot *tgbotapi.BotAPI
}

func NewTelegram(tokenVal string) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(tokenVal)
	if err != nil {
		return nil, err
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	tel := &TelegramBot{}
	tel.bot = bot

	return tel, err
}

func (t TelegramBot) Send(chatId int64, message string) error {
	msg := tgbotapi.NewMessage(chatId, message)

	m, err := t.bot.Send(msg)
	if err != nil {
		log.Fatalln("Telegram_bot: Fail to send message!")
		return err
	}
	log.Println("Telegram_bot: Success to send message")
	log.Println(m)

	return nil
}