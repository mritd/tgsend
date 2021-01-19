package main

import (
	"fmt"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

type Telegram struct {
	bot *tb.Bot
}

func NewTelegram(api, token string) (*Telegram, error) {
	bot, err := tb.NewBot(tb.Settings{
		URL:    api,
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 5 * time.Second},
	})
	if err != nil {
		return nil, err
	} else {
		return &Telegram{bot: bot}, nil
	}
}

func (tg *Telegram) SendMessage(msg string, to int64, markdown bool) error {
	opt := &tb.SendOptions{}
	if markdown {
		opt.ParseMode = tb.ModeMarkdown
		msg = fmt.Sprintf("```%s```", msg)
	}

	_, err := tg.bot.Send(tb.ChatID(to), msg, opt)
	return err
}

func (tg *Telegram) SendFile(filePath, fileName, mime, caption string, to int64) error {
	_, err := tg.bot.Send(tb.ChatID(to), &tb.Document{
		File:     tb.FromDisk(filePath),
		Caption:  caption,
		MIME:     mime,
		FileName: fileName,
	})
	return err
}

func (tg *Telegram) SendImage(imagePath, caption string, to int64) error {
	_, err := tg.bot.Send(tb.ChatID(to), &tb.Photo{
		File:    tb.FromDisk(imagePath),
		Caption: caption,
	})
	return err
}
