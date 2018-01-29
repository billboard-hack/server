package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"os/user"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// message represents a single message
type message map[string]interface{}

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
			if name, ok := m["name"]; ok {
				log.Printf("recv from %s: %v", name, m["message"])
			} else {
				log.Printf("recv: %v\n", m)
			}
		}
	}()

	forward := make(chan message)
	go func() {
		forward <- message{"name": *name, "message": "join!!"}
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			v := make(message)
			t := sc.Text()
			if strings.HasPrefix(t, "{") && strings.HasSuffix(t, "}") {
				err := json.Unmarshal([]byte(t), &v)
				if err != nil {
					v = message{"error": err.Error()}
				}
			} else {
				v["name"] = *name
				v["message"] = t
			}
			forward <- v
		}
	}()

	for {
		select {
		case v := <-forward:
			err := c.WriteJSON(v)
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
