package main

import (
	"flag"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type message struct {
	//ToAddr   string `json:"-"`
	//FromAddr string `json:"-"`
	FromUser string
	ToUser   string
	Ping     bool
	Pong     bool
	Message  string
}

func main() {

	addr := flag.String("addr", "127.0.0.1", "ip интерфейса")
	port := flag.Int("port", 9900, "номер порта")
	user := flag.String("user", "anonim", "имя пользователя")
	flag.Parse()

	addressBook := map[string]*net.TCPConn{} //< Имя пользователя и соединение с ним

	incoming := make(chan message) //< Очередь входящих сообщений
	outgoing := make(chan message) //< Очередь исходящих сообщений

	go listener(*addr, *port, addressBook, incoming, outgoing)
	go sender(addressBook, incoming, outgoing, *user)

	go writer(incoming)
	go reader(outgoing)

	insig := make(chan os.Signal)
	signal.Notify(insig, syscall.SIGINT)
	<-insig
}
