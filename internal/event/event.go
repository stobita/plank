package event

import (
	"encoding/json"
	"errors"

	"github.com/stobita/plank/internal/model"
	"github.com/stobita/plank/internal/presenter"
	"github.com/stobita/plank/internal/usecase"
)

type Broker interface {
	Run()
	AddClient(usecase.EventClient)
	RemoveClient(usecase.EventClient)
	broadcast([]byte)
	PushAddCardEvent(*model.Board) error
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

func (b *broker) broadcast(msg []byte) {
	for c := range b.clients {
		c.SendChannel <- msg
	}
}

func (b *broker) PushAddCardEvent(m *model.Board) error {
	if m == nil {
		return errors.New("Board must be set")
	}
	res, err := presenter.GetAddCardEvent(m)
	if err != nil {
		return err
	}
	json, err := json.Marshal(res)
	b.broadcast(json)
	return nil
}
