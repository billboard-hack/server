package server

import (
	"fmt"
	"log"
	"net/http"

	"billboard-wsserver/trace"

	"github.com/gorilla/websocket"
)

var store map[*message]bool

type room struct {

	// forward is a channel that holds incoming messages
	// that should be forwarded to the other clients.
	forward chan *message

	// join is a channel for clients wishing to join the room.
	join chan *client

	// leave is a channel for clients wishing to leave the room.
	leave chan *client

	// clients holds all current clients in this room.
	clients map[*client]bool

	// tracer will receive trace information of activity
	// in the room.
	tracer trace.Tracer
}

// newRoom makes a new room that is ready to go.
func newRoom() *room {
	return &room{
		forward: make(chan *message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer:  trace.Off(),
	}
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			// joining
			r.clients[client] = true
			r.tracer.Trace("New client joined")
		case client := <-r.leave:
			// leaving
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("Client left")
		case msg := <-r.forward:
			r.tracer.Trace("Message received: ", msg)
			if v, ok := (*msg)["stock"]; ok {
				fmt.Println("stock ok")
				if stock, ok := v.(bool); stock && ok {
					fmt.Println("stock true ok")
					store[msg] = true
				}
			}
			// forward message to all clients
			for client := range r.clients {
				client.send <- msg
				if v, ok := (*msg)["communication"]; ok {
					fmt.Println("communication ok")
					if s, ok := v.(string); ok && s == "broadcast" {
						fmt.Println("broadcast ok")
						for msg := range store {
							client.send <- msg
						}
					}
				}
				r.tracer.Trace(" -- sent to client")
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.tracer.Trace(req.URL.RawQuery)
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	client := &client{
		socket: socket,
		send:   make(chan *message, messageBufferSize),
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
