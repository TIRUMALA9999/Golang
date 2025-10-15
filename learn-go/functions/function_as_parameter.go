package main
import "fmt"

func compute(f func(int,int) int, a int, b int) int{
	return f(a,b)
}

func main(){
	result:= compute(func(a,b int) int{return a*b},3,4)
	fmt.Println(result)
}