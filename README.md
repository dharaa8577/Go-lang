Assignment - 6 
6.1
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
package main

import "fmt"

type person struct {
	firstname        string
	lastname         string
	favouritecountry string
}

func main() {
	var Dhara person

	Dhara.firstname = "Dhara"
	Dhara.lastname = "Goswami"
	Dhara.favouritecountry = "India"
	fmt.Println(Dhara)
}
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
6.2
package main

import "fmt"

func main() {

	var lastname map[string]string

	fmt.Println("lastname", lastname)

	persontype := map[string]string{
		"Firstname":        "Dhara",
		"Lastnmae":         "Goswami",
		"Favouritecountry": "India",
	}

	fmt.Println("persontype", persontype)

}
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
6.3.1
package main

import "fmt"

type vehicle struct {
	Doors int
	color string
}

func main() {
	var Accord vehicle

	Accord.Doors = 4
	Accord.color = "Grey"

	fmt.Println(Accord)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
6.3.2
package main

import "fmt"

type truck struct {
	truck     string
	Fourwheel bool
	color     string
}

func main() {
	ford := truck{

		truck:     "Sixdoor",
		Fourwheel: false,
		color:     "Red",
	}
	fmt.Println(ford)
}
