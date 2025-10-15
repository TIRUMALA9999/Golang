package main
import "fmt"
func main(){
	var n int
	var count int
	fmt.Print("Enter a natural number:")
	fmt.Scanln(&n)

	for i:=1;i<=n;i++{
		count=count+i
	}
	fmt.Println("sum of natural numbers is:",count)
}