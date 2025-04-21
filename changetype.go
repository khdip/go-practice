package main

import (
	"fmt"
	"strconv"
)

type Giventype interface {
	changeType() string
}

type Intvalue struct {
	intValue int64
}

type Floatvalue struct {
	floatValue float64
}

type Values struct {
	i Intvalue
	f Floatvalue
}

func (v Values) changeType() (string, string) {
	x := strconv.Itoa(int(v.i.intValue))
	y := strconv.FormatFloat(v.f.floatValue, 'f', 4, 64)
	return x, y
}

func convertToString(gt Values) {
	fmt.Println(gt.changeType())
}

func main() {
	a := Intvalue{20}
	b := Floatvalue{10.08}
	myValue := Values{a, b}
	convertToString(myValue)
}
