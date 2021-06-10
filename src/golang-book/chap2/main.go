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
//	fmt.Println("Hello World")
//
//	fmt.Println("1 + 1 = ", 1+1)
//	fmt.Println("1 + 1 =", 1.0+1.0)
//
//	fmt.Println(len("Hello World"))
//	fmt.Println("Hello World"[1])
//	fmt.Println("Hello " + "World")
//
//	fmt.Println(true && true)
//	fmt.Println(true && false)
//	fmt.Println(true || true)
//	fmt.Println(true || false)
//	fmt.Println(!true)
//
//	var x string = "Hello World"
//	fmt.Println(x)
//	var y string
//	x = "Hello World"
//	fmt.Println(y)
//
//	/*
//		The type is not necessary because the Go compiler
//		is able to infer the type based on the literal value you
//		assign the variable.
//
//		Generally you should use this shorter form whenever
//		possible.
//	*/
//	z := "Hello World assigned"
//	fmt.Println(z)
//	num := 5
//	fmt.Println(num)
//
//	fmt.Println(global)
//
//	const constant string = "I am a constant"
//	fmt.Println(constant)
//	const anotherConstant = "I am a constant, but I do not require explicit type mentioning"
//	fmt.Println(anotherConstant)
//
//	var (
//		a = "We"
//		b = "are"
//		c = "multiple"
//	)
//	fmt.Println(a, " ", b, "", c)
//	const (
//		we        = "we"
//		are       = "are"
//		constants = "constants"
//	)
//	fmt.Println(we, " ", are, "", constants)
//
//	fmt.Println(`On
//multiple
//lines`)
//
//	i := 1
//	for i <= 10 {
//		fmt.Println(i)
//		i++
//	}
//	for i := 1; i <= 10; i++ {
//		fmt.Println(i)
//	}
//
//	//fmt.Println("on ") fmt.Println("same line")
//	fmt.Println("on ")
//	fmt.Println("same line")
//
//	for j := 1; j <= 10; j++ {
//		if (j % 2) == 0 {
//			fmt.Println("even") //if an else should always be indented as done here
//		} else {
//			fmt.Println("odd")
//		}
//	}
//
//	val := 4
//
//	switch val {
//	case 0:
//		fmt.Println("Zero")
//	case 1:
//		fmt.Println("One")
//	case 2:
//		fmt.Println("Two")
//	case 3:
//		fmt.Println("Three")
//	case 4:
//		fmt.Println("Four")
//		fallthrough
//	case 5:
//		fmt.Println("Five")
//	default:
//		fmt.Println("Unknown Number")
//	}
//
//	var arr [5]int
//	arr[4] = 100
//	fmt.Println(arr)
//
//	//Go compiler won't allow you to create variables
//	//that you never use. Since we don't use i inside of our
//	//loop
//	//A single _ (underscore) is used to tell the compiler that
//	//we don't need this.
//	//var total float64 = 0
//	//for i, value := range x {
//	//	total += value
//	//} fmt.Println(total /
//	//	float64(len(x)))
//
//	var total int = 0
//	for _, value := range arr {
//		total += value
//	}
//	fmt.Println("total: ", total)
//	fmt.Println(total / len(arr))
//
//	arr2 := [5]float64{98, 93, 77, 82, 83}
//	fmt.Println(arr2)
//
//	arr3 := [5]float64{
//		98,
//		93,
//		77,
//		82,
//		//83,
//	}
//	fmt.Println(arr3)
//	//	a major issue with arrays:
//	//their length is fixed and part of the array's type name.
//
//	/*A slice is a segment of an array. Like arrays slices are
//	indexable and have a length. Unlike arrays this length
//	is allowed to change.*/
//
//	var slice []float64
//	fmt.Println(slice)
//	slice2 := make([]float64, 5)
//	fmt.Println(slice2)
//
//	/*Slices are always associated
//	with some array, and although they can never be
//	longer than the array, they can be smaller.*/
//	sliceWithCapacity := make([]float64, 5, 10)
//	fmt.Println(sliceWithCapacity)
//
//	iAmAnArray := [5]float64{1,2,3,4,5}
//	iAmASlice := iAmAnArray[1:5]
//
//	fmt.Println(iAmAnArray)
//	fmt.Println(iAmASlice)

	slice_1 := []int{1,2,3}
	slice_2 := append(slice_1, 4, 5)
	fmt.Println(slice_1, slice_2)

	slice1 := []int{1,2,3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

	//var iAmAMap map[string]int // iAmAMap is a map of strings to ints
	//iAmAMap["key"] = 10
	//fmt.Println(iAmAMap)
	/*The problem with our program is that maps have to be
	initialized before they can be used. We should have
	written this:*/

	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	anotherMap := make(map[int]int)
	anotherMap[1] = 10
	fmt.Println(anotherMap[1])
	delete(anotherMap, 1)
	fmt.Println(anotherMap)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])
	fmt.Println(elements["aloy"])
	/*Accessing an element of a map can return two values
	instead of just one. The first value is the result of the
	lookup, the second tells us whether or not the lookup
	was successful.*/
	name, ok := elements["Un"]
	fmt.Println(name, ok)

	/*
	shorter way to create a map
	elements := map[string]string{
		"H": "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B": "Boron",
		"C": "Carbon",
		"N": "Nitrogen",
		"O": "Oxygen",
		"F": "Fluorine",
		"Ne": "Neon",
	}*/

}
