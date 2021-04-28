package websocket

import (
	"api_rest_sexual/internal/interfaces"
	"api_rest_sexual/internal/models"
	"github.com/antoniodipinto/ikisocket"
)

func OnlineUsers(kws *ikisocket.Websocket) error {
	kws.SetAttribute("type", "online")
	return nil
}

func ImplementOnlineUsers(online models.Online) error {
	_, err := interfaces.GetOnline(online.Identifier)
	if err != nil {
		interfaces.CreateOnline(online)
	}

	return nil
}

func ChatsWS(kws *ikisocket.Websocket) error {
	kws.SetAttribute("type", "chats")
	return nil
}