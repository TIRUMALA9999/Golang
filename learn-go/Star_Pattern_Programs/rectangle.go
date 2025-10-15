package main

import "fmt"

func main(){
	var n int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&n)
	for i:=1;i<=4;i++{
		for j:=1;j<=n;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}