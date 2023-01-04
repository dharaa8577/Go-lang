package main

import "fmt"

func main() {
	var x string
	fmt.Println(" enter the input")
	fmt.Scan(&x)
	if x == "golangtutorial" {
		fmt.Println("welcome")
	} else {
		fmt.Println("end")
	}

}
