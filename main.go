package main

import "fmt"

type HacerSonidoDeAnimal interface {
	sonido() string
	datos() (string, int)
}

type Perro struct {
	name string
	age  int
}

type Gato struct {
	name string
	age  int
}

func (P Perro) sonido() string {
	return "woof"
}
func (G Gato) sonido() string {
	return "miau"
}

func (P Perro) datos() (string, int) {
	return P.name, P.age
}

func (G Gato) datos() (string, int) {
	return G.name, G.age
}

func Imprimir(H HacerSonidoDeAnimal) {
	fmt.Println(H.sonido())
	fmt.Println("-----------------------------")
	name, age := H.datos()
	fmt.Printf("%s tiene %d años\n\n", name, age)
}

func main() {

	perro := Perro{
		name: "Mordiscos",
		age:  3,
	}

	var gato Gato
	gato = Gato{
		name: "Bigotes",
		age:  1,
	}

	Imprimir(gato)
	Imprimir(perro)

}
