package pack

import (
	"fmt"
)

type Math interface {
	Add() int
	Multiply() int
}

type Calc struct {
	Left, Right int
}

func (c *Calc) Add() int {
	return c.Left + c.Right
}

func (c *Calc) Multiply() int {
	return c.Left * c.Right
}

//There is no explicit declaration requires that Calc implements Math interface.
//Calc simply need to define the methods provided by interface.

func DoCalc(math Math) {
	sum := math.Add()
	fmt.Println(sum)
}

func Demo1() {
	calc := Calc{Left: 1, Right: 2}
	//calc.Left = 1
	//calc.Right = 2
	DoCalc(&calc)
}
