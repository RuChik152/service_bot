package bot

import (
	"belivr_service_bot/db"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var BOT *tgbotapi.BotAPI
var BOT_CHANEL = make(chan []byte, 5)

func InitBOT() {
	BOT = bot()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := BOT.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch {
		case regexp.MustCompile(`^version_\d+$`).MatchString(update.Message.Command()):
			go func(chatID int64, version string) {

				var commitData db.CommitData

				id := chatID

				if commit_id, err := strconv.Atoi(version); err == nil {
					db.GetCommitData(commit_id, "commits", &commitData)

					newMessage := tgbotapi.NewMessage(id, fmt.Sprintf("------------------------------\nВерсия сборки: 0.%d\nАвтор: %s\nКоментарий:\n%s\n------------------------------\nsha: %s", commitData.ID, commitData.AUTHOR, commitData.COMMENT, commitData.SHA))
					if _, err = BOT.Send(newMessage); err != nil {
						log.Panic("Ошибка отправки сообщения боту: ", err)
					}

				}
			}(update.Message.Chat.ID, update.Message.Command()[8:])
		// case "sayhi":
		// 	msg.Text = "Hi :)"
		// case "status":
		// 	msg.Text = "I'm ok."
		case update.Message.Command() == "echo":
			go func(chatID int64) {
				id := chatID
				log.Println("<<Сработал echo>> : CHAI_ID: ", id)
				go EchoBot(BOT_CHANEL, id)
			}(update.Message.Chat.ID)
		default:
			msg.Text = "I don't know that command"
		}
	}

}

func bot() *tgbotapi.BotAPI {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic("Error init bot: ", err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	log.Println("Success bot init")
	return bot
}

func EchoBot(ch chan []byte, chat_id int64) {
	for chunk := range ch {
		if msg, err := defineQuery(chunk); err != nil {
			log.Println("EchoBot: ", err)
		} else {
			log.Println("Отправляю запрос боту: ", msg)
			if len(msg) <= 4000 {
				newMessage := tgbotapi.NewMessage(chat_id, msg)
				if _, err := BOT.Send(newMessage); err != nil {
					log.Panic("Ошибка отправки сообщения боту: ", err)
				}
			} else {
				shortMsg := msg[:4000]
				newMessage := tgbotapi.NewMessage(chat_id, shortMsg)
				if _, err := BOT.Send(newMessage); err != nil {
					log.Panic("Ошибка отправки сообщения боту: ", err)
				}
			}

		}
	}
}
