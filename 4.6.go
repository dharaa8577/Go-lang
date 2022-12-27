// Addition
package main
import (
    "fmt"
)
 
func add(v *int) {
    *v = *v + 10
 
    fmt.Println(*v)
}
 
func main() {
    a := 10
    add(&a)
    //fmt.Println(a)
 
}
 
 
// Substraction
package main
 
import (
    "fmt"
)
 
func Sub(v *int) {
    *v = *v - 10
 
    fmt.Println(*v)
}
 
func main() {
    a := 10
    Sub(&a)
    //fmt.Println(a)
 
}
 
// Multiplication
package main
 
import (
    "fmt"
)
 
func Mul(v *int) {
    *v = *v * 10
 
    fmt.Println(*v)
}
 
func main() {
    a := 10
    Mul(&a)
    //fmt.Println(a)
 
}
// Divison
package main
 
import (
    "fmt"
)
 
func Div(v *int) {
    *v = *v / 10
 
    fmt.Println(*v)
}
 
func main() {
    a := 10
    Div(&a)
    //fmt.Println(a)
 
}
 
 
 
 
 
 
 
 
