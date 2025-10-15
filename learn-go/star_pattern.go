package main

import "fmt"

func main() {
    var n int
    fmt.Print("Enter a natural number: ")
    fmt.Scanln(&n)

    // Outer loop controls the number of rows
    for i := 1; i <= n; i++ {
        // Inner loop prints stars for each row
        for j := 1; j <= i; j++ {
            fmt.Print("*")
        }
        // Move to the next line after each row
        fmt.Println()
    }
}
