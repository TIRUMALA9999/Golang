package main
import "fmt"

func sum(arr [5]int) int{
	total:=0
	for _,nums:=range arr{
		total=total+nums
	}
	return total
}

func main(){
	fmt.Println(sum([5]int{1,2,3,4,5}))
}