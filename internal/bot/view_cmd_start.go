package bot

import (
	"context"
	"errors"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"news-feed-bot/internal/botkit"
)

func ViewCmdStart() botkit.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		if update.Message == nil {
			return errors.New("update.Message is nil")
		}

		chatID := update.Message.Chat.ID
		userName := update.Message.From.FirstName

		welcomeMessage := tgbotapi.NewMessage(chatID, "👋 HI, "+userName+"!\n")

		if _, err := bot.Send(welcomeMessage); err != nil {
			return err
		}

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		newsBot := botkit.New(bot)
		newsBot.RegisterCmdView("start", ViewCmdStart())

		if err := newsBot.Run(ctx); err != nil {
			if errors.Is(err, context.Canceled) {
				log.Printf("[ERROR] failed to start bot: %v", err)
				return err
			}
		}
		log.Println("bot stopped")

		return nil
	}
}
