package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options
var clients = make(map[*websocket.Conn]bool)

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
	defer c.Close()
	clients[c] = true
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

func StartHTTPServer() {
	http.HandleFunc("/ws", ws)
	http.HandleFunc("/echo", echo)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
	fmt.Println("Starting HTTP Server @ http://localhost:8080")
}
