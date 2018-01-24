// Author: "Shun Yokota"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package server

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"billboard-server/trace"
)

// getPortFromEnv gets port number form environment variable.
// if environment variable for port number is not set, returns defaultPort of the argument.
func getPortFromEnv(defaultPort string) string {
	depEnv := os.Getenv("DEPLOY_ENV")
	switch depEnv {
	case "heroku":
		return os.Getenv("PORT")
	default:
		if port := os.Getenv("UMAYADO_PORT"); len(port) > 0 {
			return port
		}
		return defaultPort
	}
}

// Start begins this system
func Start() error {
	port := flag.String("port", getPortFromEnv("8080"), "The host of the application.")

	flag.Parse() // parse the flags

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "billboard websocket server\n") })
	http.Handle("/ws", r)

	// get the room going
	go r.run()

	// start the web server
	log.Println("Starting web server on", *port)
	return http.ListenAndServe(":"+*port, nil)

}
