package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gobuffalo/packr/v2"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options
var clients = make(map[*websocket.Conn]bool)
var router *mux.Router

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func ws(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade: ", err)
		return
	}
	log.Println("Websocket client connected")
	defer c.Close()
	clients[c] = true
	UpdateScoreBoard(&s)
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
	}
}

func UpdateScoreBoard(s *Scoreboard) {
	for client := range clients {
		err := client.WriteJSON(s)
		if err != nil {
			log.Printf("Websocket error: %s", err)
			client.Close()
			delete(clients, client)
		}
	}
}

func UpdateTimer(function string) {
	for client := range clients {
		message := []byte(function)
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("Websocket error: %s", err)
			client.Close()
			delete(clients, client)
		}
	}
}

func StartHTTPServer() {
	r := mux.NewRouter()
	r.HandleFunc("/ws", ws)
	r.HandleFunc("/echo", echo)
	r.Handle("/logo", http.FileServer(http.Dir(logoPath)))
	// box := packr.New("static", "./static")
	// r.Handle("/", http.FileServer(box))
	router = r
	go func() {
		log.Fatal(http.ListenAndServe(":8080", r))
	}()
	fmt.Println("Starting HTTP Server @ http://localhost:8080")
}

func UpdateStaticRoute() {
	defaultTheme := packr.New("static", "./static")
	bglTheme := packr.New("bgl", "./themes/bgl")
	switch selectedTheme {
	case "default":
		router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(defaultTheme)))
	case "bgl":
		router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(bglTheme)))
	default:
		router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir(filepath.Join(themePath, selectedTheme)))))
	}
}
