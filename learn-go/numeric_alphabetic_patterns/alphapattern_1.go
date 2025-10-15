package main
import "fmt"

func main(){
	var n int
	var p int =65
	fmt.Println("Enter the number: ")
	fmt.Scanln(&n)
	//fmt.Println("Enter the p value: ")
	//fmt.Scanln(&p)
	for i:=1;i<=n;i++{
		for j:=1;j<=i;j++{
			fmt.Print(string(p))
		}
		p+=1
		fmt.Println()
	}
}