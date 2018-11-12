package array

import "fmt"

//imprime diferentes modos de hacer un array
func Array() {
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

//imprime maneras de hacer slice
func Slice() {
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
