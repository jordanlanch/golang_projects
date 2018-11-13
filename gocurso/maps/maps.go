package maps

//GetMap develve un map con llave tipo string y valores en valores tipo int
func GetMap(name string) int {
	mapTest := map[string]int{
		"Juan":   18,
		"Yohan":  24,
		"Freddy": 31,
	}
	mapTest["llave1"] = 1
	mapTest["llave2"] = 100

	delete(mapTest, "llave1")
	return mapTest[name]
}
