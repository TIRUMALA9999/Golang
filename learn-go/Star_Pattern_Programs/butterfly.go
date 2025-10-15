package main
import "fmt"

func main(){
	var n int
	fmt.Println("Enter the  number: ")
	fmt.Scanln(&n)
	for i:=1;i<=n;i++{
		for j:=1;j<=i;j++{
			fmt.Print("*")
		}
		for j:=1;j<=2*(n-i);j++{
			fmt.Print(" ")
		}
		for j:=1;j<=i;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i:=n;i>=0;i--{
		for j:=1;j<=i;j++{
			fmt.Print("*")
		}
		for j:=1;j<=2*(n-i);j++{
			fmt.Print(" ")
		}
		for j:=1;j<=i;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}