package main

import (
	"errors"
	"fmt"
)

func main() {
	var number int64 = 1000000000

	// If not specified, default is the architecutre of the operating system
	// int64 will be used
	var number2 int = 1000000000

	var unsigned_int uint = 1000

	//alias for int32 = rune
	var number3 rune = 12456

	//alias for uint8 = byte
	var number4 byte = 123

	var realNumber float32 = 123.45
	var realNumber2 float64 = 12355555.45

	// Table ASC number for B = 66
	char := 'B'

	// Every go data type has its zero value
	var text string

	var erro error = errors.New("Internal error")

	fmt.Println(number, number2, unsigned_int, number3, number4, realNumber, realNumber2, char, text, erro)
}