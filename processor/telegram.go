package processor

import (
	"fmt"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (p *Processor) StartTG() {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := p.bot.GetUpdatesChan(u)

	log.Print("Bot started")

	for u := range updates {
		if u.Message == nil {
			continue
		}

		err := p.handleMsg(u.Message)
		if err != nil {
			log.Printf("can't handle messages: %s", err)
		}
	}
}

func (p *Processor) handleMsg(Message *tgbotapi.Message) error {

	if !validateMsg(Message.Text) {
		return fmt.Errorf("url is not valid: %s", Message.Text)
	}

	err := p.LoadContent(Message.Text)
	if err != nil {
		return fmt.Errorf("can't load content: %w", err)
	}

	msg := tgbotapi.NewMessage(Message.Chat.ID, "Контент отправлен в канал")
	p.bot.Send(msg)

	files, err := p.content()
	if err != nil {
		return fmt.Errorf("can't get content: %w", err)
	}

	err = p.sendContent(files)
	if err != nil {
		return fmt.Errorf("can't send content: %w", err)
	}

	return nil
}

func (p *Processor) sendContent(files []*os.File) error {

	var media []interface{}

	var img int
	var vid int

	for _, tf := range files {

		if strings.HasSuffix(tf.Name(), ".jpg") || strings.HasSuffix(tf.Name(), ".jpeg") || strings.HasSuffix(tf.Name(), ".png") {
			img++
		}

		if strings.HasSuffix(tf.Name(), ".mp4") {
			vid++
		}

	}

	for _, tf := range files {

		if (strings.HasSuffix(tf.Name(), ".jpg") || strings.HasSuffix(tf.Name(), ".jpeg") || strings.HasSuffix(tf.Name(), ".png")) && img >= 2 {
			media = append(media, tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(tf.Name())))
		}

		if strings.HasSuffix(tf.Name(), ".mp4") {
			media = append(media, tgbotapi.NewInputMediaVideo(tgbotapi.FilePath(tf.Name())))
		}

	}

	msg := tgbotapi.NewMediaGroup(p.channelID, media)
	_, err := p.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("can't send media group: %w", err)
	}

	return nil

}

func validateMsg(msg string) bool {
	return strings.HasPrefix(msg, "https://instagram.com/") || strings.HasPrefix(msg, "https://www.instagram.com/")
}
