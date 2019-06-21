package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

// sender - Отправляет сообщения из очереди собеседникам
func sender(addressBook map[string]*net.TCPConn, incoming chan message, outgoing chan message, user string) {
	for {
		msg := <-outgoing

		// Если true - сообщение служебноя для установления связи
		isPing := false

		// Ищем получателя в адресной книге
		conn, exists := addressBook[msg.ToUser]
		if !exists {

			// Если пользователь ещё неизвестен - устанавливаем с ним соединение полагая что пользователь ввёл адрес вместо имени
			tcpaddr, err := net.ResolveTCPAddr("tcp", msg.ToUser)
			if err != nil {
				fmt.Println("пользователь не найден, укажите адрес")
				continue
			}

			c, err := net.DialTCP("tcp", nil, tcpaddr)
			if err != nil {
				fmt.Printf("! не могу соедениться c: %s\n", msg.ToUser)
				continue
			}

			go listenConnection(c, addressBook, incoming, outgoing)

			isPing = true
			conn = c

		}

		// Если сообщение служеное
		msg.Ping = isPing

		// Добавляем в сообщение своё имя
		msg.FromUser = user

		// Сериализуем сообщение
		outbytes, err := json.Marshal(msg)
		if err != nil {
			log.Fatalf("не могу маршалить сообщение: %s", err)
		}

		if _, err := conn.Write(outbytes); err != nil {
			fmt.Printf("! пользователь недоступен: %s", conn.RemoteAddr())
			delete(addressBook, conn.RemoteAddr().String())
		}

	}
}
