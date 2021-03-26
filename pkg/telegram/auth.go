package telegram

import (
	"context"
	"fmt"
)

func (b *Bot) generateAuthorizationLink(chatId int64) (string, error) {
	redirectURL := b.generateRedirectUrl(chatId)
	requestToken, err := b.pocketClient.GetRequestToken(context.Background(), redirectURL)
	if err != nil {
		return "", err
	}

	return b.pocketClient.GetAuthorizationURL(requestToken, redirectURL)
}

func (b *Bot) generateRedirectUrl(chatId int64) string {
	return fmt.Sprintf("%s?chat_id=%d", b.redirectURl, chatId)
}
