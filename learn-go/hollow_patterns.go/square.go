package main
import "fmt"

func main(){
	var n int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&n)
	for i:=1;i<=n;i++{
		for j:=1;j<=n;j++{
			if i==1 || i==n || j==1 || j==n{
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}