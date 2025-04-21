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

func (i Intvalue) changeType() string {
	return strconv.Itoa(int(i.intValue))
}

func (f Floatvalue) changeType() string {
	return strconv.FormatFloat(f.floatValue, 'f', 4, 64)
}

func typeChange(gt Giventype) string {
	return gt.changeType()
}

func main() {
	myValue1 := Intvalue{20}
	myValue2 := Floatvalue{10.80}
	fmt.Printf("%s\n", typeChange(myValue1))
	fmt.Printf("%s\n", typeChange(myValue2))
}
