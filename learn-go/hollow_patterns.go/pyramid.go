package main
import "fmt"

func main() {
    var n int
    fmt.Print("Enter the number of rows: ")
    fmt.Scanln(&n)

    for i := 1; i <= n; i++ {
        // Spaces before stars
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }
        // Hollow part
        for j := 1; j <= 2*i-1; j++ {
            if j == 1 || j == 2*i-1 || i==n {
                fmt.Print("*")
            } else {
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
}
