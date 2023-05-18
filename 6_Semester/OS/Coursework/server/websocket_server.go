package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	ws   *websocket.Conn
	hub  *Hub
	send chan interface{}
}

type GeneratorClient struct {
	Client
}

type ProcessorClient struct {
	Client
}

type Hub struct {
	chats      map[int]map[*Client]bool
	clients    map[*Client]bool
	processors map[*ProcessorClient]bool
	generators map[*GeneratorClient]bool
	// send       chan *dto.Message
	register   chan *Client
	unregister chan *Client
}

func CreateHub() *Hub {
	return &Hub{
		chats:   make(map[int]map[*Client]bool),
		clients: make(map[*Client]bool),
		// send:       make(chan *dto.Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (hub *Hub) removeClient(client *Client) {
	if _, ok := hub.chats[client.chat_id]; !ok {
		return
	}

	if _, ok := hub.chats[client.chat_id][client]; !ok {
		return
	}

	delete(hub.chats[client.chat_id], client)
	delete(hub.clients, client)
	close(client.send)

	if len(hub.chats[client.chat_id]) == 0 {
		delete(hub.chats, client.chat_id)
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register:
			_, ok := hub.chats[client.chat_id]
			if !ok {
				hub.chats[client.chat_id] = make(map[*Client]bool)
			}

			hub.chats[client.chat_id][client] = true
			hub.clients[client] = true

		case client := <-hub.unregister:
			hub.removeClient(client)

		case message := <-hub.send:
			for client, _ := range hub.chats[message.ChatID] {
				select {
				case client.send <- message:
					fmt.Println("send to client message ", message)
				default:
					hub.removeClient(client)
				}
			}
			// chat, err := models.GetChatByID(message.ChatID)
			// if err != nil {
			// 	fmt.Println(err)
			// 	continue
			// }

			// users, err := models.GetChatParticipants(message.ChatID)
			// if err != nil {
			// 	fmt.Println(err)
			// 	continue
			// }
		}
	}
}

func AddClient(ws *websocket.Conn, hub *Hub) *Client {

	return &Client{
		ws: ws,
	}
}

func AddGeneratorClient(ws *websocket.Conn, hub *Hub) *GeneratorClient {
	return &GeneratorClient{
		Client: *AddClient(ws, hub),
	}
}

func AddProcessorClient(ws *websocket.Conn, hub *Hub) *ProcessorClient {
	return &ProcessorClient{
		Client: *AddClient(ws, hub),
	}
}
