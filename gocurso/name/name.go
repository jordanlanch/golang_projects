package name

import "fmt"

// GetName obtiene y retorna el nombre de la persona
func GetName() string {
	var name string
	name = "Sin Nombre"
	fmt.Print("Ingresa tu nombre:")
	fmt.Scanf("%s\n", &name)
	return name
}

//GetLastName obtiene y retorna el apellido de la persona
func GetLastName() string {
	var lastname string
	lastname = "Sin Apellido"
	fmt.Print("Ingresa tu Apellido:")
	fmt.Scanf("%s\n", &lastname)
	return lastname
}
