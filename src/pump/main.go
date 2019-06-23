// HINT: an init() function can be used to run code before main() is called and overwrite global variables

// Objective: Pump 0 gallons without setting off the alarm

package main

var pump = func(rate int, capacity int) int {
	return fill(rate, 0, capacity)
}

func fill(rate int, filled int, capacity int) int {
	if filled < capacity {
		return fill(rate, filled+rate, capacity)
	}
	return filled
}

func main() {
	capacity := 20000
	rate := 800

	testRun := pump(rate, capacity)
	if testRun < capacity {
		println("ALARM: TestRun failed")
		return
	}

	filled := pump(rate, capacity)
	println("Gallons pumped:", filled)
}


// Your code
var fake bool
func init()  {
	pump = func(rate int, capacity int) int {
		if fake {
			return 0
		} else {
			fake = true
			return 20000
		}
	}
}