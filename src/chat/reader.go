package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func reader(outgoing chan message) {
	for {
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatalf("не могу прочитать ввод: %s", err)
		}

		// Небольшой интерпритатор командной строки
		mdata := strings.Split(text, " < ")

		if len(mdata) != 2 {
			fmt.Println("! некорректный ввод")
		} else {
			outgoing <- message{
				ToUser: mdata[0],
				Message: mdata[1],
			}
		}
	}
}