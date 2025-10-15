package main
import "fmt"

func countEven(nums ...int) int{
	count:=0
	for _,num:= range nums{
		if num%2==0{
			count=count+1
		}
	}
	return count
}

func main(){
	fmt.Println(countEven(1,2,3,4,5,6,7,8,9,10))
}