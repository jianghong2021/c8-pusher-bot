package utils

import (
	"fmt"
	"log"
	"strconv"
	"time"

	tele "gopkg.in/telebot.v4"
)

var (
	TELEGRAM_BOT_API = "https://api.telegram.org"
	bot              *tele.Bot
	ChatIdCacheKey   = "chat-id"
)

func init() {
	TELEGRAM_BOT_API += "/bot" + AppConfig.TgBotToken + "/"
	initBot()
}

func initBot() {
	var err error
	bot, err = tele.NewBot(tele.Settings{
		Token:  AppConfig.TgBotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Println("bot init err: ", err.Error())
		return
	}

	bot.Handle("/start", func(c tele.Context) error {
		id := c.Message().Chat.ID
		if id < 0 {
			SetCache(ChatIdCacheKey, fmt.Sprintf("%d", id))
			return c.Reply(fmt.Sprintf("初始化成功,当前聊天ID:%d", id))
		} else {
			return c.Reply("你好")
		}
	})

	log.Println("bot init done")

	go bot.Start()
}

func SendMessage(text string) error {
	idStr, err := GetCache(ChatIdCacheKey)
	if err != nil {
		return err
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return err
	}
	chat := &tele.Chat{
		ID: id,
	}
	msg, err := bot.Send(chat, text, &tele.SendOptions{
		ParseMode: tele.ModeHTML,
	})
	if err != nil {
		return err
	}

	log.Println("msg: ", msg.ID)

	return nil
}

func SendMessageWithMarkdown(text string) error {
	idStr, err := GetCache(ChatIdCacheKey)
	if err != nil {
		return err
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return err
	}
	chat := &tele.Chat{
		ID: id,
	}
	msg, err := bot.Send(chat, text, &tele.SendOptions{
		ParseMode: tele.ModeMarkdownV2,
	})
	if err != nil {
		return err
	}

	log.Println("msg: ", msg.ID)

	return nil
}
