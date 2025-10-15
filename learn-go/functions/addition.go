package main
import "fmt"

func add(a int, b int) int{
	return a+b
}

func main(){
	var a int
	var b int
	fmt.Println("Enter a value: ")
	fmt.Scanln(&a)
	fmt.Println("Enter b value: ")
	fmt.Scanln(&b)
	fmt.Println(add(a,b))
}