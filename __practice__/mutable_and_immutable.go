package main

import "fmt"

func boolMutability() {
	x := true
	y := x

	x = false

	fmt.Printf("\nBool mutibility\n")
	fmt.Printf("x: (%T) %-5v %p \n", x, x, &x)
	fmt.Printf("y: (%T) %-5v %p \n", y, y, &y)
}
func intMutability() {
	x := 5
	y := x

	x = 7

	fmt.Printf("\nInt mutibility\n")
	fmt.Printf("x: (%T) %v %p \n", x, x, &x)
	fmt.Printf("y: (%T) %v %p \n", y, y, &y)
}

func floatMutability() {
	x := float64(5)
	y := x

	x = 7

	fmt.Printf("\nFloat mutibility\n")
	fmt.Printf("x: (%T) %v %p \n", x, x, &x)
	fmt.Printf("y: (%T) %v %p \n", y, y, &y)
}

func strMutability() {
	x := "sara"
	y := x

	x = "reza"

	fmt.Printf("\nStr mutibility\n")
	fmt.Printf("x: (%T) %v %p \n", x, x, &x)
	fmt.Printf("y: (%T) %v %p \n", y, y, &y)
}

func sliceMutability() {
	x := []int{1, 2, 3}
	y := x

	x[0] = 100

	fmt.Printf("\nSlice mutibility\n")
	fmt.Printf("x: (%T) %v %p \n", x, x, &x)
	fmt.Printf("y: (%T) %v %p \n", y, y, &y)
	fmt.Printf("actual data address: %p \n", &x[0])
	fmt.Printf("actual data address: %p \n", &y[0])
}

func arrayMutability() {
	x := [3]int{1, 2, 3}
	y := x

	x[0] = 5

	fmt.Printf("\nArray mutibility\n")
	fmt.Printf("x: (%T) %v %p \n", x, x, &x)
	fmt.Printf("y: (%T) %v %p \n", y, y, &y)
}

func main() {
	boolMutability()
	intMutability()
	floatMutability()
	strMutability()
	sliceMutability()
	arrayMutability()
}
