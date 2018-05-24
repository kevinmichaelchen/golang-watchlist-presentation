package code

import "fmt"

func main() {
	p := &Person{Steps: 5}
	p.AnnounceSteps()
	p.Move()
	p.AnnounceSteps()
}

type Person struct {
	Steps int
}

func (p *Person) Move() {
	p.Steps += 5
}

func (p *Person) AnnounceSteps() {
	fmt.Printf("I am %d inches tall\n", p.Steps)
}