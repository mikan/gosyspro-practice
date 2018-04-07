package main

import "fmt"

// Talker interface
type Talker interface {
	Talk()
}

// Greeter struct
type Greeter struct {
	name string
}

// Talk prints name with message
func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
	g.name = "updated"
}

func main() {
	var talker Talker
	talker = &Greeter{"wozozo"}
	talker.Talk() // Hello, my name is wozozo

	g := &Greeter{"wozozo"}
	g.Talk() // Hello, my name is wozozo
	g.name = "foo"
	g.Talk() // Hello, my name is foo

	g2 := Greeter{"wozozo"}
	g2.Talk() // Hello, my name is wozozo
	g2.Talk() // Hello, my name is wozozo
}
