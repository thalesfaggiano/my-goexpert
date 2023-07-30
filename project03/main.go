package main

import "fmt"

func main() {
	myArray := []int{10, 20, 30, 40}
	//myArray[0] = 10
	//myArray[1] = 20
	//myArray[2] = 30

	fmt.Printf("len = %d, cap = %d, array = %d\n", len(myArray[2:]), cap(myArray[2:]), myArray[2:])
}
