package main

import (
	"fmt"
	"strings"
)

//main cambia la ultima palabra a Mayúsculas
func main() {
	var str string
	fmt.Println("Escriba una palabra")
	fmt.Scanf("%s", &str)

	fmt.Println(changeLastWord(str))
}

func changeLastWord(str string) string {
	//palabra --> str
	//palabr --> str[0:len(str)-1]   A -->strings.ToUpper(str[len(str)-1:]
	return str[0:len(str)-1] + strings.ToUpper(str[len(str)-1:])
}
