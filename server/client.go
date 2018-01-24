package server

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// client represents a single chatting user.
type client struct {

	// socket is the web socket for this client.
	socket *websocket.Conn

	// send is a channel on which messages are sent.
	send chan *message

	// room is the room this client is chatting in.
	room *room

	// name holds user name
	name string
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		var msg message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("read error: %v\n", err)
			return
		}
		c.room.forward <- &msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			fmt.Printf("write error: %v\n", err)
			return
		}
	}
}
