package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("new incoming websocket connection from ", ws.RemoteAddr())
	s.conns[ws] = true
	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error :", err)
			continue
		}
		ws.Write(buf[:n])
		msg := buf[:n]
		fmt.Println(string(msg))
		ws.Write([]byte("Thank you for the msg!!!"))
	}
}

func main() {
	s := NewServer()
	http.Handle("/ws", websocket.Handler(s.handleWS))
	http.ListenAndServe(":3000", nil)
}
