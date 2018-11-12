package flow

import (
	"fmt"
	"strings"
)

//IfTest recibe un entero y tetorna si es par o impar
func IfTest(number int) {
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

//StringLibrary manejo de Strings
func StringLibrary() {
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

//SwitchTest manejo de sitch en condicionales
func SwitchTest() {
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
