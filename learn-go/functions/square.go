package main
import "fmt"

func square(a int) int{
	return a*a
}

func main(){
	value:=square(5)
	fmt.Println(value)
}