package main

import "fmt"

func main() {
    slice := []int{10, 20, 30,40,50,60}
    fmt.Println(len(slice)) // 3
    fmt.Println(cap(slice)) // 3

    slice = append(slice, 40)
    fmt.Println(slice)      // [10 20 30 40]
    fmt.Println(len(slice)) // 4
    fmt.Println(cap(slice)) // capacity may increase (usually doubled)
}
