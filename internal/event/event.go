package event

import (
	"github.com/stobita/plank/internal/usecase"
)

type Broker interface {
	Run()
	AddClient(usecase.EventClient)
	RemoveClient(usecase.EventClient)
	Broadcast([]byte)
}

type broker struct {
	clients             map[*Client]bool
	AddClientChannel    chan *Client
	RemoveClientChannel chan *Client
}

type Client struct {
	SendChannel chan []byte
}

func NewBroker() Broker {
	return &broker{
		clients:             make(map[*Client]bool),
		AddClientChannel:    make(chan *Client),
		RemoveClientChannel: make(chan *Client),
	}
}

func (b *broker) Run() {
	go b.run()
}

func (b *broker) run() {
	for {
		select {
		case client := <-b.AddClientChannel:
			b.clients[client] = true
		case client := <-b.RemoveClientChannel:
			if _, ok := b.clients[client]; ok {
				delete(b.clients, client)
			}
		}
	}
}

func (b *broker) AddClient(c usecase.EventClient) {
	b.AddClientChannel <- c.(*Client)
}

func (b *broker) RemoveClient(c usecase.EventClient) {
	b.RemoveClientChannel <- c.(*Client)
}

func (b *broker) Broadcast(msg []byte) {
	for c := range b.clients {
		c.SendChannel <- msg
	}
}
