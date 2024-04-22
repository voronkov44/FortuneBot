package main

import (
	"math/rand"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const TOKEN = "YOUR TOKEN"

var bot *tgbotapi.BotAPI

var chatId int64

var fortuneTellerNames = [3]string{"ванга", "нострадамус", "понсар"}

var answers = []string{
	"Да",
	"Нет",
	"Возможно, но я не уверен.",
	"Мне кажется, что это зависит от обстоятельств.",
	"Надо подумать над этим вопросом.",
	"Это сложный вопрос, требующий глубокого анализа.",
	"Я склоняюсь к тому, что можно найти ответ на этот вопрос.",
	"Возможно, стоит обратиться за консультацией к эксперту.",
	"Может быть, но нужно учитывать различные факторы.",
	"Я бы предпочел не делать предположений.",
	"Сложно сказать однозначно, необходимо провести исследование.",
	"Пожалуй, стоит обсудить этот вопрос более подробно.",
}

func connectWithTelegram() {
	var err error
	if bot, err = tgbotapi.NewBotAPI(TOKEN); err != nil {
		panic("Cannot connect to Telegram")
	}
}

func sendMessage(msg string) {
	msgConfig := tgbotapi.NewMessage(chatId, msg)
	bot.Send(msgConfig)
}

func isMessageForFortuneTeller(update *tgbotapi.Update) bool {
	if update.Message == nil || update.Message.Text == "" {
		return false
	}

	msgInLowerCase := strings.ToLower(update.Message.Text)
	for _, name := range fortuneTellerNames {
		if strings.Contains(msgInLowerCase, name) {
			return true
		}
	}
	return false
}

func getFortuneTellersAnswer() string {
	index := rand.Intn(len(answers))
	return answers[index]
}

func sendAnswer(update *tgbotapi.Update) { //отвечает на сообщение которое в апдейте
	msg := tgbotapi.NewMessage(chatId, getFortuneTellersAnswer())
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func main() {
	connectWithTelegram()

	updateConfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message != nil && update.Message.Text == "/start" {
			chatId = update.Message.Chat.ID
			sendMessage(" Привет👋, задай мне вопрос назвав мое имя (Ванга, Нострадамус, Понсар). " +
				"\nОтветом на вопрос должны быть \"да\" либо \"нет\". " +
				"\nПример: Ванга, мне сегодня стоит выйти на улицу? ")
		}
		if isMessageForFortuneTeller(&update) { //ответ от бота
			sendAnswer(&update)
		}
	}
}
