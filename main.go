package main

import (
	"os"
	"os/exec"
    "log"
    "fmt"
    "strings"
    "gopkg.in/telegram-bot-api.v4"
)

func main() {
    bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
    if err != nil {
        log.Panic(err)
    }

    bot.Debug = true

    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates, err := bot.GetUpdatesChan(u)

    for update := range updates {
        text := update.Message.Text

        if strings.HasPrefix(text, "/run") {
        	splitter := strings.Split(text, " ")
        	args := []string{}
        	if len(splitter) == 0 {
        		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error on command execution")
        		bot.Send(msg)
        	}
        	cmd := splitter[1]
        	if len(splitter) >2 {
        		args = splitter[2:]
        	}
        	fmt.Println(args)
        	cmdOut := exec.Command(cmd, args...)
        	cmdOut.Wait()
        	res,err := cmdOut.Output()
        	if err != nil {
        		fmt.Printf("%v", err)
        	}

        	msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(res))
        	bot.Send(msg)

        }

        if strings.HasPrefix(text, "script") {

        }
    }
}