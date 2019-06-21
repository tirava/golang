package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

func listenConnection(conn *net.TCPConn, addressBook map[string]*net.TCPConn, incoming chan message, outgoing chan message) {
	for {
		buff := make([]byte, 2048)
		bufflen, err := conn.Read(buff)
		if err != nil {
			// Поиск имени сохранённого пользователя
			name := ""
			for n, c := range addressBook {
				if conn == c {
					name = n
				}
			}

			fmt.Printf("! пользователь отключился: %s\n", name)
			delete(addressBook, name)
			break
		}

		msg := message{}
		if err := json.Unmarshal(buff[:bufflen], &msg); err != nil {
			log.Fatalf("не могу распарсить сообщение: %s", err)
		}

		// Если сообщения не служебные, то ставим их в пользовательскую очередь сообщений
		if !msg.Ping && !msg.Pong {
			incoming <- msg
		}

		if msg.Pong {
			fmt.Printf("! принял соединение пользователь: %s\n", msg.FromUser)
		}

		// Если служебное сообщение пинг, отвечаем понгом
		if msg.Ping {
			fmt.Printf("! принято соединение от: %s\n", msg.FromUser)
			addressBook[msg.FromUser] = conn
			m := message{
				Pong: true,
				ToUser: msg.FromUser,
			}
			outgoing <- m
		}

		addressBook[msg.FromUser] = conn

	}
}