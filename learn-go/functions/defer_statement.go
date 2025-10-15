//Defer Statement (defer function call until surrounding function returns)
package main
import "fmt"

func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
// Output:
// hello
// world
