package main
import "fmt"

func main(){
	var n int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&n)
	for i:=0;i<=n-1;i++{
		for j:=1;j<=n-i;j++{
			fmt.Print(" ")
		}
		for j:=1;j<=2*i+1;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i:=n-2;i>=0;i--{
		for j:=1;j<=n-i;j++{
			fmt.Print(" ")
		}
		for j:=1;j<=2*i+1;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}	
