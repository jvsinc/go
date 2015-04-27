package main

import "fmt"

type Person struct {
	name string
}
type talker interface {
	Talk() (string, string, int)
}

func (p Person) Talk() (string, string, int) {
	fmt.Println("Hi name is ", p.name)
	return p.name, "Jimmy", 1
}

func main() {

	p := Person{"John"}
	fmt.Println(p.name)

	name, middleName, num := p.Talk()
	fmt.Println("Name is: ", name)
	fmt.Println("Middle Name is: ", middleName, num)

	makeHimTalk(p)

}

func makeHimTalk(t talker) {
	t.Talk()
}
