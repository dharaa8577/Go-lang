package main

import "fmt"

func main() {

	var states [4]string

	states[0] = "Alaska"
	states[1] = "Oregon"
	states[2] = "California"
	states[3] = "Florida"

	fmt.Println("states initial: ", states)
	updateThirdElement(&states[2])
	fmt.Println("states updated: ", states)

}

func updateThirdElement(states *string) {
	*states = "Texas"
	fmt.Println(*states)
}
