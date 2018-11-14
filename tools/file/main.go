package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Traer información de la ejecución del programa
	_, filename, _, ok := runtime.Caller(0)
	// Valida si hay ejecución, si no existe ejecición envia un panic
	if !ok {
		panic("No caller information")
	}
	/* Leemos el archivo obteniendo la ruta del archivo a leer,
	en este caso el archivo esta en la misma carpeta que el archivo main.go
	*/
	data, err := ioutil.ReadFile(path.Dir(filename) + "/archivo.txt")
	// Si hay error detenemos la ejecución
	check(err)
	// Imprimimos en pantalla el contenido del archivo.
	fmt.Println(string(data))
}
