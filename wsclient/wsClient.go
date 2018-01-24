package main

import (
	"bufio"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"os/user"
	"time"

	"github.com/gorilla/websocket"
)

// message represents a single message
type message struct {
	Name    string `json:"name,omitempty"`
	Message string `json:"message,omitempty"`
}

func main() {
	addr := flag.String("addr", "localhost:8080", "http service address")
	user, err := user.Current()
	if err != nil {
		return
	}
	name := flag.String("name", user.Username, "user name")

	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws", RawQuery: "name=" + *name}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer c.Close()
		defer close(done)
		for {
			var m message
			err := c.ReadJSON(&m)
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv from %s: %s", m.Name, m.Message)
		}
	}()

	forward := make(chan string)
	go func() {
		forward <- "join!!"
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			forward <- sc.Text()
		}
	}()

	for {
		select {
		case t := <-forward:
			err := c.WriteJSON(&message{Name: *name, Message: t})
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			c.Close()
			return
		}
	}
}
