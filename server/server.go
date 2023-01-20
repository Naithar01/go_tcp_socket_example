package main

import (
	"log"
	"net"
)

type Server struct {
	Server       net.Listener
	ListenerPort string
}

func InitServer(ListenerPort string) *Server {
	return &Server{
		ListenerPort: ListenerPort,
	}
}

func (s *Server) StartServer() error {
	server, err := net.Listen("tcp", s.ListenerPort)

	if err != nil {
		log.Println(err.Error())
	}

	s.Server = server

	defer server.Close()

	log.Println("Running Server_____")

	s.AcceptClient()

	return nil
}

func (s *Server) AcceptClient() {
	for {
		connect_client, err := s.Server.Accept()

		if err != nil {
			log.Println("Accept Client Error: " + err.Error())

			continue
		}

		defer connect_client.Close()

		log.Println("Connect Success Client")

		go s.RecvComment(connect_client)
	}
}

func (s *Server) RecvComment(connect_client net.Conn) {
	recvBuf := make([]byte, 2048)

	for {
		recv, err := connect_client.Read(recvBuf)

		if err != nil {
			log.Println("Recv Error: " + err.Error())

			break
		}

		if recv > 0 {
			msg := recvBuf[:recv]

			log.Println(string(msg))

			connect_client.Write([]byte(msg))
		}

	}
}

func main() {
	server := InitServer(":8000")

	server.StartServer()
}
