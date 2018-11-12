package main

import (
	"fmt"

	"github.com/jordanlanch/projects/array"
	"github.com/jordanlanch/projects/flow"
	"github.com/jordanlanch/projects/name"
	"github.com/jordanlanch/projects/numbers"
	"github.com/jordanlanch/projects/structs"
)

const helloWorld string = "Hola %s %s, Bienvenido al fascinante mundo de Go\n"

func main() {
	firstname := name.GetName()
	lastname := name.GetLastName()
	var (
		a = 1
		b = "dos"
		c = false
	)
	fmt.Printf(helloWorld, firstname, lastname)
	fmt.Printf("%b %s %t\n", a, b, c)
	numbers.GetFloatMax()
	countWords()
	array.Array()
	array.Slice()
	var number int
	fmt.Println("Ingrese un número del 1 al 10: ")
	fmt.Scanf("%d\n", &number)
	structs.Guess()
	structs.ThreeFors()
	flow.IfTest(number)
	flow.StringLibrary()
	flow.SwitchTest()
}

func countWords() {
	fmt.Println(string("Hola"[3]))
	fmt.Println("la cantidad de palabras de hola es ", len("hola"))
}
