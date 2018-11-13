package structs

import (
	"fmt"
	"math/rand"
	"time"
)

//Guess compara el resultado del usuario con un numero aleatorio
func Guess() {
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

//ThreeFors 3 tipos de hacer for
func ThreeFors() {
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

//PlatziCourse -- STRUCTS --  struct
type PlatziCourse struct {
	Name   string
	Slug   string // url del curso
	Skills []string
}

//PlatziCareer struct
type PlatziCareer struct {
	Name   string
	Slug   string // url de la carrera
	Skills []string
}

//Suscribe metodo
func (p PlatziCourse) Suscribe(name string) {
	fmt.Printf("El usuario %s, se ha inscripto al curso de %s\n", name, p.Name)
}

//Suscribe metodo
func (p PlatziCareer) Suscribe(name string) {
	fmt.Printf("El usuario %s, se ha inscripto a la carrera de %s\n", name, p.Name)
}
