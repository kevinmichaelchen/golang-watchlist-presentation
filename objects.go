package main

import "fmt"

func main() {
	p := &Person{Height: 5}
	p.AnnounceHeight()
	p.GrowTaller()
	p.AnnounceHeight()
}

type Person struct {
	Height int
}

func (p *Person) GrowTaller() {
	p.Height += 5
}

func (p *Person) AnnounceHeight() {
	fmt.Printf("I am %d inches tall\n", p.Height)
}