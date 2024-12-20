package bot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"news-feed-bot/internal/botkit"
)

func ViewCmdStart() botkit.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		// Отправка сообщения "Hello, world!"
		if _, err := bot.Send(tgbotapi.NewMessage(update.FromChat().ID, "Hello, world!")); err != nil {
			return err
		}

		return nil
	}
}
