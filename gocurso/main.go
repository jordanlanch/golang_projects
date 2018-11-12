package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

const helloWorld string = "Hola %s %s, Bienvenido al fascinante mundo de Go\n"

func main() {
	name := getName()
	lastname := getLastName()
	var (
		a = 1
		b = "dos"
		c = false
	)
	fmt.Printf(helloWorld, name, lastname)
	fmt.Printf("%b %s %t\n", a, b, c)
	//getFloatMax()
	//countWords()
	array()
	slice()
	var number int
	fmt.Println("Ingrese un número del 1 al 10: ")
	fmt.Scanf("%d\n", &number)
	ifTest(number)
	//guess()
	//threeFors()
	stringLibrary()
	switchTest()
}
func getName() string {
	var name string
	name = "Sin Nombre"
	fmt.Print("Ingresa tu nombre:")
	fmt.Scanf("%s\n", &name)
	return name
}

func getLastName() string {
	var lastname string
	lastname = "Sin Apellido"
	fmt.Print("Ingresa tu Apellido:")
	fmt.Scanf("%s\n", &lastname)
	return lastname
}
func getFloatMax() {
	minInt8 := math.MinInt8
	maxInt8 := math.MaxInt8
	minInt16 := math.MinInt16
	maxInt16 := math.MaxInt16
	minInt32 := math.MinInt32
	maxInt32 := math.MaxInt32
	minInt64 := math.MinInt64
	maxInt64 := math.MaxInt64
	maxUnit8 := math.MaxUint8
	maxUnit16 := math.MaxUint16
	maxUnit32 := math.MaxUint32
	maxFloat32 := math.MaxFloat32
	maxFloat64 := math.MaxFloat64

	fmt.Printf("Range of Int8 = %d to %d.\n", minInt8, maxInt8)
	fmt.Printf("Range of Int16 = %d to %d.\n", minInt16, maxInt16)
	fmt.Printf("Range of Int32 = %d to %d.\n", minInt32, maxInt32)
	fmt.Printf("Range of Int64 = %d to %d.\n", minInt64, maxInt64)

	fmt.Printf("MaxUint8 = %d.\n", maxUnit8)
	fmt.Printf("MaxUint16 = %d.\n", maxUnit16)
	fmt.Printf("MaxUint32 = %d.\n", maxUnit32)

	fmt.Printf("Max Float32 = %f.\n", maxFloat32)
	fmt.Println("Max Float64 =", maxFloat64)
}
func countWords() {
	fmt.Println(string("Hola"[3]))
	fmt.Println("la cantidad de palabras de hola es ", len("hola"))
}
func array() {
	// vat name[size] type
	var array [3]string
	fmt.Println(array)    // [ ]
	fmt.Println(array[0]) // 'Nothing'
	// Initialization
	array[0] = "First"
	array[1] = "Second"
	array[2] = "Thrid"
	fmt.Println(array)
	fmt.Println(array[0]) // First

	array[2] = "Last"
	fmt.Println("array forma larga", array) // [First Second Last]

	arr2 := [3]int{1, 2, 3}
	fmt.Println("array forma corta", arr2)
}

func slice() {
	// var name []type
	var slice []string
	fmt.Println(slice) // []
	// Initialization
	slice = append(slice, "First")
	fmt.Println(slice[0]) // First
	slice = append(slice, "Second")
	fmt.Println(slice) // [First Second]

	slice[0] = "One"
	fmt.Println(slice) // [One Second]
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
func guess() {
	var number int
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	secret := random.Intn(10)
	// fmt.Println(secret)

	for i := 0; i < 10; i++ {
		fmt.Print("Guess de number (1 - 10): ")
		fmt.Scanf("%d", &number)
		if number == secret {
			fmt.Println("Correct!")
			break
		} else {
			fmt.Println("Worng!")
		}
	}
}

func threeFors() {
	// 1. Normal version
	fmt.Println("For, normal version")
	for i := 0; i < 100; i += 10 {
		fmt.Println(i)
	}
	// 2. While version
	fmt.Println("For, while version")
	i := 100
	for i > 0 {
		i -= 10
		fmt.Println(i)
	}
	// 3. Infinite version
	fmt.Println("For, infinite version")
	j := 0
	for {
		j++
		fmt.Println(j)
		if j == 10 {
			fmt.Println("The end!")
			break
		}
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
