package main
import "fmt"
func main(){
	var n int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&n)
	for i:=1;i<=n;i++{
		for j:=1;j<=i;j++{
			if j==1 || i==j{
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		for j:=1;j<=2*(n-i);j++{
			fmt.Print(" ")
		}
		for j:=1;j<=i;j++{
			if j==i || j==1{
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	for i:=n-1;i>=1;i--{
		for j:=1;j<=i;j++{
			if j==1 || i==j{
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		for j:=1;j<=2*(n-i);j++{
			fmt.Print(" ")
		}
		for j:=1;j<=i;j++{
			if j==i || j==1{
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}