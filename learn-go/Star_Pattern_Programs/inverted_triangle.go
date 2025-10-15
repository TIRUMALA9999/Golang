package main
import "fmt"

func main(){
	var n int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&n)
	for i:=1;i<=n;i++{
		for j:=n;j>=i;j--{
			fmt.Print("*")
		}
		fmt.Println()
	}
}