package main

import "fmt"

func main() {
	// Explicitly declaration of variable
	var variable1 string = "variable1"
	// go infer the variable type
	variable2 := "variable2"

	var (
		variable3 string = "variable3"
		variable4 string = "variable4"
	)

	variable5, variable6 := "variable5", "variable6"

	// In go we can switch values from variables without an auxiliar one.
	variable5, variable6 = variable6, variable5

	fmt.Println(variable1, variable2, variable3, variable4, variable5, variable6)
}