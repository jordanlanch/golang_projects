package structs

import (
	"fmt"
)

//Platzi para PlatziCourse y PlatziCareer
type Platzi interface {
	Suscribe(name string)
}

//deferTest utilizado comunmente para cerrar archivos
func deferTest() {
	fmt.Println("La función interfaceTest ha Terminado")
}

//InterfaceTest funcion que llama a los struct y llama a la interfaz con un unico nombre
func InterfaceTest() {
	defer deferTest()
	platziCourse := PlatziCourse{Name: "Go", Slug: "go", Skills: []string{"backend"}}
	platziCareer := new(PlatziCareer)
	platziCareer.Name = "Carrera Go"
	platziCareer.Slug = "go"
	callSuscribe(platziCareer)
	callSuscribe(platziCourse)
}

func callSuscribe(p Platzi) {
	p.Suscribe("Jordan")
}
