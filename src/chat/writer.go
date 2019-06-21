package main

import "fmt"

func writer(incoming chan message) {
	for {
		msg := <- incoming
		fmt.Printf("%s > %s", msg.FromUser, msg.Message)
	}
}