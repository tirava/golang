// HINT: startup() accepts an interface, so login() can return a struct that matches what validSequence() is expecting

// Objective: Start the Jet Ski

package main

import (
	"reflect"
)

func main() {
	println("Logging in...")
	authorized := startup(login())
	if reflect.ValueOf(authorized).Bool() {
		println("Starting the engine")
		return
	}
	println("Startup failed")
}

func validSequence(i int, el interface{}) bool {
	//fmt.Println(reflect.TypeOf(el).String())
	//fmt.Println(!reflect.ValueOf(el).IsNil())
	//fmt.Println(reflect.ValueOf(el).Elem().NumField())
	//fmt.Println(reflect.TypeOf(reflect.ValueOf(el).Elem().Field(0).Interface()).String())
	//fmt.Println(int(reflect.ValueOf(el).Elem().Field(0).Int()), i)
	//fmt.Println(!reflect.ValueOf(reflect.ValueOf(el).Elem().Field(1).Interface()).IsNil())
	//return true

	return reflect.TypeOf(el).String() == "*main.Sequence" &&
		!reflect.ValueOf(el).IsNil() &&
		reflect.ValueOf(el).Elem().NumField() == 2 &&
		reflect.TypeOf(reflect.ValueOf(el).Elem().Field(0).Interface()).String() == "int" &&
		int(reflect.ValueOf(el).Elem().Field(0).Int()) == i*i-i &&
		!reflect.ValueOf(reflect.ValueOf(el).Elem().Field(1).Interface()).IsNil()
}

func startup(seq interface{}) bool {
	for i := 0; i < 5; i++ {
		if !validSequence(i, seq) {
			return false
		}
		seq = reflect.ValueOf(seq).Elem().Field(1).Interface()
	}

	return true
}


type Sequence struct {
	Num int
	Seq *Sequence
}

func login() *Sequence {
	s := Sequence{
		Num: 0,
		Seq: new(Sequence),
	}
	*s.Seq = Sequence{
		Num: 0,
		Seq: new(Sequence),
	}
	*s.Seq.Seq = Sequence{
		Num: 2,
		Seq: new(Sequence),
	}
	*s.Seq.Seq.Seq = Sequence{
		Num: 6,
		Seq: new(Sequence),
	}
	*s.Seq.Seq.Seq.Seq = Sequence{
		Num: 12,
		Seq: new(Sequence),
	}

	return &s
}