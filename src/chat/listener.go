package main

import (
	"fmt"
	"log"
	"net"
)

// listener - Слушает интерфейс, принимает входящие соединения и заносит имена в адресную книгу
func listener(addr string, port int, addressBook map[string]*net.TCPConn, incoming chan message, outgoing chan message) {
	tcpaddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		log.Fatalf("не могу резолвить адрес: %s", err)
	}

	// Слушаем входящие к нам
	listener, err := net.ListenTCP("tcp", tcpaddr)
	if err != nil {
		log.Fatalf("не могу встать на интерфейс: %s", err)
	}

	fmt.Printf("! слушаю соединения на: %s\n", listener.Addr())

	// Принимаем входящие запросы на подключение
	for {

		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatalf("не могу принять соединение: %s", err)
		}

		go listenConnection(conn, addressBook, incoming, outgoing)

	}
}