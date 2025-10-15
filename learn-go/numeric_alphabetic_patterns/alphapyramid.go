package main
import "fmt"

func main(){
	var n int
	var p int=65
	fmt.Println("Enter the number: ")
	fmt.Scanln(&n)
	fmt.Println("Enter p value: ")
	fmt.Scanln(&p)
	for i:=1;i<=n;i++{
		for j:=1;j<=n-i;j++{
			fmt.Print(" ")
		}
		for j:=1;j<=2*i-1;j++{
			fmt.Print(string(p))
		}
		p+=1
		fmt.Println()

	}
}