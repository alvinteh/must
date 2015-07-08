package must

import (
	"strconv"
)

//GetBA returns the first argument as a byte array
func GetBA(a ...interface{}) []byte {
	return a[0].([]byte)
}

//GetF32 returns the first argument as a float32
func GetF32(a ...interface{}) float32 {
	return a[0].(float32)
}

//GetF64 returns the first argument as a float64
func GetF64(a ...interface{}) float64 {
	return a[0].(float64)
}

//GetF returns the first argument as a float64
func GetF(a ...interface{}) float64 {
	return GetF64(a)
}

//GetI returns the first argument as an int
func GetI(a ...interface{}) int {
	return a[0].(int)
}

//GetS returns the first argument as a string
func GetS(a ...interface{}) string {
	return a[0].(string)
}

//F32 returns a float32 representation of the argument
func F32(a interface{}) float32 {
	f, err := strconv.ParseFloat(a.(string), 32)

	if err != nil {
		panic(err)
	}

	return float32(f)
}

//F64 returns a float64 representation of the argument
func F64(a interface{}) float64 {
	f, err := strconv.ParseFloat(a.(string), 64)

	if err != nil {
		panic(err)
	}

	return f
}

//BA returns a byte array representation of the argument
func BA(a interface{}) []byte {
	return []byte(a.(string))
}

//F returns a float64 representation of the argument
func F(a interface{}) float64 {
	return F64(a)
}

//I returns an int representation of the argument
func I(a interface{}) int {
	i, err := strconv.Atoi(a.(string))

	if err != nil {
		panic(err)
	}

	return i
}

//S returns a string representation of the argument
func S(a interface{}) string {
	return a.(string)
}
