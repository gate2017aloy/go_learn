package main

import "fmt"

//Go is lexically scoped using blocks
//variable exists within the nearest curly braces { }
var global string = "My name is aloy"

//func main()  {
func main() {
	// this is a comment
	/*
	   	go run main.go
	   	The go run command takes the subsequent files (separated
	      by spaces), compiles them into an executable
	      saved in a temporary directory and then runs the program.
	*/

	// Our very own hello world program
	fmt.Println("Hello World")

	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("1 + 1 =", 1.0+1.0)

	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	var x string = "Hello World"
	fmt.Println(x)
	var y string
	x = "Hello World"
	fmt.Println(y)

	/*
		The type is not necessary because the Go compiler
		is able to infer the type based on the literal value you
		assign the variable.

		Generally you should use this shorter form whenever
		possible.
	*/
	z := "Hello World assigned"
	fmt.Println(z)
	num := 5
	fmt.Println(num)

	fmt.Println(global)

	const constant string = "I am a constant"
	fmt.Println(constant)
	const anotherConstant = "I am a constant, but I do not require explicit type mentioning"
	fmt.Println(anotherConstant)

	var (
		a = "We"
		b = "are"
		c = "multiple"
	)
	fmt.Println(a, " ", b, "", c)
	const (
		we        = "we"
		are       = "are"
		constants = "constants"
	)
	fmt.Println(we, " ", are, "", constants)

	fmt.Println(`On 
multiple 
lines`)

	i :=1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//fmt.Println("on ") fmt.Println("same line")
	fmt.Println("on "); fmt.Println("same line")

	for j := 1; j <= 10; j++ {
		if(j % 2) == 0 {
			fmt.Println("even")  //if an else should always be indented as done here
		} else {
			fmt.Println("odd")
		}
	}




}
