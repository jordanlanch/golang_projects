package main

import (
	"fmt"
	"time"
	/*"github.com/jordanlanch/projects/gocurso/array"
	"github.com/jordanlanch/projects/gocurso/flow"
	"github.com/jordanlanch/projects/gocurso/name"
	"github.com/jordanlanch/projects/gocurso/numbers"
	"github.com/jordanlanch/projects/gocurso/structs"*/)

const helloWorld string = "Hola %s %s, Bienvenido al fascinante mundo de Go\n"

func main() {
	/*firstname := name.GetName()
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
	flow.SwitchTest()*/
	// fmt.Println(maps.GetMap("Juan"))
	/*cursos := Course{Name: "Go", Slug: "go", Skills: []string{"1", "2"}}
	fmt.Println(cursos)

	pythonCourse := new(Course)
	pythonCourse.Name = "Curso Python"
	pythonCourse.Slug = "curso-python"
	pythonCourse.Skills = []string{"Data Science", "Backend"}
	fmt.Println(pythonCourse)*/
	//structs.InterfaceTest()
	/*number, err := numbers.Sum(50, 50)
	if err != nil {
		panic(err)
	}
	fmt.Println(number)*/
	//pointerTest()
	go forGo(500)
	go forGo(400)
	time.Sleep(1000 * time.Millisecond)
}

func countWords() {
	fmt.Println(string("Hola"[3]))
	fmt.Println("la cantidad de palabras de hola es ", len("hola"))
}

func pointerTest() {
	a := 100
	var b *int
	b = &a
	fmt.Println("Sin modificar")
	fmt.Println(a, *b)
	fmt.Println(&a, b)
	pointerModify(b)
	fmt.Println("Con una modificación")
	fmt.Println(a, *b)
	fmt.Println(&a, b)
}

func pointerModify(c *int) {
	*c = 10
}
func helloGo(index int) {
	fmt.Println("Hola soy un print en la Go routine #", index)
}

func forGo(n int) {
	for i := 0; i < n; i++ {
		go helloGo(i)
	}
}
