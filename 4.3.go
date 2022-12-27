package main

import "fmt"

func main() {

	num := make([]int, 10)
	num[0] = 42
	num[1] = 43
	num[2] = 44
	num[3] = 45
	num[4] = 46
	num[5] = 47
	num[6] = 48
	num[7] = 49
	num[8] = 50
	num[9] = 51

	fmt.Println("num[0:5]", num[0:5])
	fmt.Println("num[5:10]", num[5:10])
	fmt.Println("num[2:7]", num[2:7])
	fmt.Println("num[1:6]", num[1:6])
}
