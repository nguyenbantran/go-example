package basic

import "fmt"

type AnimalQuark interface {
	quark()
}

type animal struct {
}

type dog struct {
	animal
}

func (a *animal) quark() {
	fmt.Println("Animal Quark")
}

func (d *dog) quark() {
	fmt.Println("Dog quark")
}
