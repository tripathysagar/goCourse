package main

import "fmt"

type Pos struct {
	x int
}

func (p *Pos) move(delta int) {
	fmt.Printf("obj moved by %d \n", p.x+delta)
}

type Object struct {
	Pos
	test int
}

//func (o *Object) move(delta int) {
//f	fmt.Printf("obj jumed by %d \n", o.Pos.x+delta)
//f	fmt.Printf("value of the test is %d\n", o.test)
//f}

func main() {
	p := Pos{
		x: 10,
	}

	p.move(10)

	o := Object{
		Pos: Pos{
			x: 100,
		},
	}

	o.move(9)
	o.Pos.move(12)

}
