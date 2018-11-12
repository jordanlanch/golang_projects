package main

import (
	"fmt"
	"strings"

	"github.com/jordanlanch/projects/array"
	"github.com/jordanlanch/projects/bucle"
	"github.com/jordanlanch/projects/name"
	"github.com/jordanlanch/projects/numbers"
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
	//countWords()
	array.Array()
	array.Slice()
	var number int
	fmt.Println("Ingrese un número del 1 al 10: ")
	fmt.Scanf("%d\n", &number)
	ifTest(number)
	bucle.Guess()
	bucle.ThreeFors()
	stringLibrary()
	switchTest()
}

func countWords() {
	fmt.Println(string("Hola"[3]))
	fmt.Println("la cantidad de palabras de hola es ", len("hola"))
}

func ifTest(number int) {
	if number > 0 || number < 10 {
		if number%2 == 0 {
			fmt.Println("El número es par")
		} else {
			fmt.Println("El número es impar")
		}
	} else {
		fmt.Println("ERROR | El número ingresado está fuera del rango solicitado.")
	}

	if numer2 := 3; numer2 == 3 {
		fmt.Println("Entró al condicional")

	}
}

func stringLibrary() {
	text := "Hello world. Hello Platzi. Hello Go"
	// To uppercase
	fmt.Println(strings.ToUpper(text))
	// To lowercase
	fmt.Println(strings.ToLower(text))
	// Replace(string-to-modified, old-string, new-string, times-to-replace)
	fmt.Println(strings.Replace(text, "Hello", "Hi", -1))
	// Replace only first
	fmt.Println(strings.Replace(text, "Hello", "Hi", 1))
	// string => slice
	fmt.Println(strings.Split(text, "."))
}

func switchTest() {
	var number int
	fmt.Print("Type a number: ")
	fmt.Scanf("%d", &number)

	switch number {
	case 1:
		fmt.Println("The number is 1")
	default:
		fmt.Println("The number isn't 1")
	}

	switch {
	case number%2 == 0:
		fmt.Println("The number is pair")
	default:
		fmt.Println("The number is odd")
	}
}
