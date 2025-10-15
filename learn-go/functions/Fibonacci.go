package main
import "fmt"

func fib(n int) string{
	a:=0
	b:=1
	var temp int
	for i:=0;i<n;i++{
		if i<=1{
			fmt.Print(i," ")
		}else{
			temp=a+b
			fmt.Print(temp," ")
			a=b
			b=temp
		}		
	}
	return " "
}

func main(){
	fmt.Print(fib(10))
}