package main
import "fmt"

func main(){
	var a int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&a)

	if a%2==0{
		fmt.Println("It is even number")
	}else{
		fmt.Println("It is Odd number")
	}
}
