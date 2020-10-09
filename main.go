package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/urfave/cli/v2"
)

var (
	version   string
	buildDate string
	commitID  string
)

func main() {
	app := &cli.App{
		Name:    "tgsend",
		Usage:   "Telegram message send tool",
		Version: fmt.Sprintf("%s %s %s", version, commitID, buildDate),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "api",
				Usage:   "Telegram api address",
				EnvVars: []string{"TELEGRAM_ADDRESS"},
				Value:   "https://api.telegram.org",
			},
			&cli.StringFlag{
				Name:    "token",
				Usage:   "Telegram api token",
				EnvVars: []string{"TELEGRAM_TOKEN"},
			},
			&cli.Int64Flag{
				Name:     "id",
				Aliases:  []string{"i"},
				Usage:    "Telegram user or group ID",
				EnvVars:  []string{"TELEGRAM_SEND_ID"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "message",
				Aliases: []string{"m"},
				Usage:   "Telegram message to be sent",
				EnvVars: []string{"TELEGRAM_MESSAGE"},
			},
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "Telegram file to be sent",
				EnvVars: []string{"TELEGRAM_FILE"},
			},
			&cli.StringFlag{
				Name:    "image",
				Aliases: []string{"photo", "p"},
				Usage:   "Telegram image to be sent",
				EnvVars: []string{"TELEGRAM_IMAGE"},
			},
			&cli.BoolFlag{
				Name:    "markdown",
				Usage:   "Set the message format to markdown",
				EnvVars: []string{"TELEGRAM_MARKDOWN"},
				Value:   false,
			},
		},
		Authors: []*cli.Author{
			{
				Name:  "mritd",
				Email: "mritd@linux.com",
			},
		},
		Action: func(c *cli.Context) error {
			bot, err := NewTelegram(c.String("api"), c.String("token"))
			if err != nil {
				return err
			}
			var wg sync.WaitGroup
			wg.Add(3)
			go sendTxtMessage(c, bot, &wg)
			go sendFile(c, bot, &wg)
			go sendImage(c, bot, &wg)
			wg.Wait()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func sendTxtMessage(c *cli.Context, bot *Telegram, wg *sync.WaitGroup) {
	defer wg.Done()
	txtMsg := c.String("message")
	if txtMsg != "" {
		err := bot.SendMessage(txtMsg, c.Int64("id"), c.Bool("markdown"))
		if err != nil {
			logger.Errorf("failed to send txt message: %v", err)
		}
	}
}

func sendFile(c *cli.Context, bot *Telegram, wg *sync.WaitGroup) {
	defer wg.Done()
	filePath := c.String("file")
	if filePath != "" {
		info, err := os.Stat(filePath)
		if err != nil {
			logger.Error(err)
			return
		}
		err = bot.SendFile(filePath, info.Name(), "", info.Name(), c.Int64("id"))
		if err != nil {
			logger.Errorf("failed to send file: %v", err)
		}
	}
}

func sendImage(c *cli.Context, bot *Telegram, wg *sync.WaitGroup) {
	defer wg.Done()
	imagePath := c.String("image")
	if imagePath != "" {
		info, err := os.Stat(imagePath)
		if err != nil {
			logger.Error(err)
			return
		}
		err = bot.SendImage(imagePath, info.Name(), c.Int64("id"))
		if err != nil {
			logger.Errorf("failed to send file: %v", err)
		}
	}
}
