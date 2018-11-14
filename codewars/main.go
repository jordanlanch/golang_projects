package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		fmt.Println(XO("xo"))
		fmt.Println(XO("xxOo"))
		fmt.Println(XO("xxxm"))
		fmt.Println(XO("Oo"))
		fmt.Println(XO("ooom"))*/
	fmt.Println(ToNato("If you can read"))
}

/*
XO("ooxx") => true
XO("xooxx") => false
XO("ooxXm") => true
XO("zpzpzpp") => true // when no 'x' and 'o' is present should return true
XO("zzoo") => false
*/
//XO funcion
func XO(str string) bool {
	x := 0
	o := 0
	for n := range strings.ToLower(str) {
		if strings.ToLower(str)[n] == 'x' {
			x++
		}
		if strings.ToLower(str)[n] == 'o' {
			o++
		}
	}

	if x == o {
		return true
	} else {
		return false
	}
}

/**
var _ = Describe("Tests using hard-coded strings", func() {
  It("Should return a correctly translated string", func() {
     Expect(ToNato("If you can read")).To(Equal("India Foxtrot Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta"))
     Expect(ToNato("Did not see that coming")).To(Equal("Delta India Delta November Oscar Tango Sierra Echo Echo Tango Hotel Alfa Tango Charlie Oscar Mike India November Golf"))
     Expect(ToNato("go for it!")).To(Equal("Golf Oscar Foxtrot Oscar Romeo India Tango !"))
  })
})

**/

//ToNato
func ToNato(words string) string {

	//        0       1        2         3         4        5         6       7       8          9        10      11      12       13         14       15        16     17         18       19        20         21       22         23        24       25
	nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "X-ray", "Yankee", "Zulu"}
	abc := map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
		"f": 5,
		"g": 6,
		"h": 7,
		"i": 8,
		"j": 9,
		"k": 10,
		"l": 11,
		"m": 12,
		"n": 13,
		"o": 14,
		"p": 15,
		"q": 16,
		"r": 17,
		"s": 18,
		"t": 19,
		"u": 20,
		"v": 21,
		"w": 22,
		"x": 23,
		"y": 24,
		"z": 25,
	}
	r := ""
	normalizada := strings.Replace(strings.ToLower(words), " ", "", -1)
	for n := range normalizada {
		//palabr --> str[0:len(str)-1]   A -->strings.ToUpper(str[len(str)-1:]
		letra := normalizada[n : n+1]

		if len(normalizada)-1 == n {
			if letra != "!" {
				r += nato[abc[letra]]
			} else {
				r += "!"
			}

		} else {

			r += nato[abc[letra]] + " "
		}
	}

	return r
}

/**
func ToNato(words string) string {
  nato := []string{
    "Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot",
    "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November",
    "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor",
    "Whiskey", "Xray", "Yankee", "Zulu",
  }
  charToCharlie := map[rune]string{}
  for _, value := range nato {
    charToCharlie[rune(value[0])] = value
  }

  result := ""
  for _, letter := range words {
    if unicode.IsLetter(letter) {
      result += charToCharlie[unicode.ToUpper(letter)] + " "
    } else if unicode.IsPunct(letter) {
      result += string(letter)
    }
  }
  return strings.TrimSpace(result)
}**/
