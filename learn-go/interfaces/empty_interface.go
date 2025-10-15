package main
import "fmt"

func printAnything(v interface{}) {
    fmt.Println(v)
}

func main() {
    printAnything(10)
    printAnything("hello")
    printAnything(true)
}
