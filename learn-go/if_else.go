package main
import "fmt"

func main(){
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	if age>=18{
		fmt.Println("You are major")
	}else {
		fmt.Println("you are minor")
	}
}