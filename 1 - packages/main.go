package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Main function!")
	auxiliar.Write()

	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}
