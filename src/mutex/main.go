// HINT: The sync package includes the ability to use mutexes

// Objective: Fix the race condition

package main

import (
	"math/rand"
	"sync"
	"time"
)

var buttons = []string{"red", "blue", "green", "yellow", "purple"}

func main() {
	rand.Seed(111009)

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(buttons))

	sequence := make([]Button, 0)

	for x := 0; x < len(buttons); x++ {
		go addButton(x, &sequence, &waitGroup)
	}

	waitGroup.Wait()

	println("Sending Code Sequence...")
	for i, button := range sequence {
		println(i+1, ":", button.color)
	}
}

//type Counter struct {
//	x int
//}
//func (c *Counter) Read() (x int) {
//	c.x++
//	return c.x
//}
//
//var counter = Counter{x:-1}
var mux sync.Mutex

func setButton(x int, sequence *[]Button) {
// -----------
	mux.Lock()
	//x = counter.Read()
	println(x)
	*sequence = append(*sequence, Button{buttons[x]})
	seq := make([]Button, 0)
	sequence = &seq
	mux.Unlock()
// -----------
	newButton := Button{buttons[x]}
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	*sequence = append(*sequence, newButton)
}

func addButton(x int, sequence *[]Button, waitGroup *sync.WaitGroup) {
	setButton(x, sequence)
	waitGroup.Done()
}

// Button represents a colored button
type Button struct {
	color string
}