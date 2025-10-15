//In this code we declare functions without names and assigned to variables
package main
import "fmt"

func main(){
	add:=func(a,b int) int{
		return a+b
	}
	fmt.Println(add(5,7))
}

