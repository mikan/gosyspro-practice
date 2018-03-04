package main

import "fmt"

type Talker interface {
	Talk()
}

type Greeter struct {
	name string
}

func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}

func main() {
	var talker Talker
	talker = &Greeter{"wozozo"}
	talker.Talk()

	g := &Greeter{"wozozo"}
	g.Talk()
	g.name = "foo"
	g.Talk()

	g2 := Greeter{"wozozo"}
	g2.Talk()
	g2.name = "bar"
	g2.Talk()
}
