// Author: "Shun Yokota"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package server

import (
	"flag"
	"log"
	"net/http"
	"os"

	"billboard-server/trace"
)

var host = flag.String("host", ":8080", "The host of the application.")

// type options struct {
// 	ip         string
// 	port       string
// 	ssl        bool
// 	slackToken string
// }

// func init() {

// }

// // parseOptions sets options after parsing the environment variable and command line options.
// // If command line options are given, they override the environment variable.
// func parseOptions(fs *flag.FlagSet, args []string) options {
// 	ip := fs.String("ip", "", "input IPv4 address of this server")
// 	port := fs.String("port", getPortFromEnv("8000"), "input port number of this server")
// 	ssl := fs.Bool("ssl", false, "input ssl activation")
// 	slackToken := fs.String("slacktoken", os.Getenv("SLACK_TOKEN"), "input API token of slack")
// 	fs.Parse(args)
// 	return options{
// 		ip:         *ip,
// 		port:       *port,
// 		ssl:        *ssl,
// 		slackToken: *slackToken,
// 	}
// }

// // getPortFromEnv gets port number form environment variable.
// // if environment variable for port number is not set, returns defaultPort of the argument.
// func getPortFromEnv(defaultPort string) string {
// 	depEnv := os.Getenv("DEPLOY_ENV")
// 	switch depEnv {
// 	case "heroku":
// 		return os.Getenv("PORT")
// 	default:
// 		if port := os.Getenv("UMAYADO_PORT"); len(port) > 0 {
// 			return port
// 		}
// 		return defaultPort
// 	}
// }

// type routingHandler struct {
// 	next http.Handler
// }

// func (h routingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(r.URL.Path)
// 	segs := strings.Split(r.URL.Path, "/")
// 	fmt.Println(segs[1])
// 	if segs[1] == "api" || segs[1] == "auth" {
// 		http.NotFound(w, r)
// 		return
// 	}
// 	r.URL.Path = "/"
// 	h.next.ServeHTTP(w, r)
// }

// // NewUmayadoMux allocates and returns a new NewUmayadoMux.
// func NewUmayadoMux() *mux.Router {
// 	r := mux.NewRouter()
// 	r.Handle("/", http.FileServer(http.Dir("./client/dist")))
// 	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./client/dist/static"))))
// 	r.PathPrefix("/ws/{id:\\d+}").Handler(manager)
// 	return r
// }

// Start begins this system
func Start() error {
	flag.Parse() // parse the flags

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/room", r)

	// get the room going
	go r.run()

	// start the web server
	log.Println("Starting web server on", *host)
	return http.ListenAndServe(*host, nil)

	// 	//start log capture
	// 	op := parseOptions(flag.CommandLine, os.Args[1:])
	// 	serveMux := NewUmayadoMux()
	// 	fmt.Printf("%s番ポートでサーバーを起動\n", op.port)
	// 	var err error
	// 	if op.ssl {
	// 		err = http.ListenAndServeTLS(op.ip+":"+op.port, "server.crt", "server.key", serveMux)
	// 	} else {
	// 		err = http.ListenAndServe(op.ip+":"+op.port, serveMux)
	// 	}
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "%v\nサーバーを起動できませんでした。\nsever starts as localhost.\nlocalhost:8000にのみアクセスできます。\n", err)
	// 	}
	// 	return http.ListenAndServe("localhost:8000", serveMux)

}
