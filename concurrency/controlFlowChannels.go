package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{} // 0 bytes,channel for indicating the server to quit
	msgch  chan string
}

func newServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string, 64),
	}
}

func (s *Server) start() {
	fmt.Println("server starting")
	s.loop()
}

func (s *Server) quit() {
	s.quitch <- struct{}{} // to indicate to clise the chaneel quitch
}
func (s *Server) sendMessage(msg string) {
	s.msgch <- msg
}
func (s *Server) loop() {
mainloop:
	for {
		select { // used for itrating over channel
		case <-s.quitch:
			fmt.Println("quiting the server")
			break mainloop
		case msg := <-s.msgch:
			s.handleMessage(msg)
		default: // for not halting the server

		}

	}
	fmt.Println("the server is quited gracefully")
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("we received a message :", msg)
}
func main() {
	server := newServer()
	//go server.start()
	//
	//for i := 0; i < 40; i++ {
	//	server.sendMessage(fmt.Sprintf("Hello from main Func!! with value %d", i))
	//}
	//time.Sleep(time.Second * 5)

	go func() {
		time.Sleep(time.Second * 5)
		server.sendMessage(fmt.Sprintf("Hello from main Func!! with value %d", 1))
		time.Sleep(time.Second * 5)
		server.quit()
	}()
	server.start()
	time.Sleep(time.Second * 20)
}
