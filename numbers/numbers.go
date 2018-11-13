package numbers

import (
	"errors"
	"fmt"
	"math"
)

//GetFloatMax indica los rangos de cada variable
func GetFloatMax() {
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

//Sum  con validacion de errores
func Sum(a interface{}, b interface{}) (int, error) {
	switch a.(type) {
	case string:
		return 0, errors.New("El valor A es un string")
	}

	switch b.(type) {
	case string:
		return 0, errors.New("El valor B es un string")
	}
	return a.(int) + b.(int), nil
}
