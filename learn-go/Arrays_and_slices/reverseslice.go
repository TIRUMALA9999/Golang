package main
import "fmt"

func reverseslice(s []int) {
	for i,j:=0,len(s)-1;i<j;i,j=i+1,j-1{
		s[i],s[j]=s[j],s[i]
	}
}

func main(){
	numbers:= []int{5, 4, 3, 6, 8, 98, 345, 678, 9, 4}
	fmt.Println("Original slice: ",numbers)

	reverseslice(numbers)
	fmt.Println("Reversed slice: ",numbers)
}