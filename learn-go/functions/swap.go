// Multiple Parameters and Return Values
/* 1. Here we can have multiple parameters
   2. Functions can return multiple values */

package main
import "fmt"

func swap(a,b string) (string,string){
	return b,a
}

func main(){
	x,y:=swap("2","3")
	fmt.Print(x,y)
}