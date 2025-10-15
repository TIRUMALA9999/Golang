package main
import "fmt"

func sliceavg(slice []int) float64{
	total:=0
	for _,nums:= range slice{
		total=total+nums
	}
	return float64(total)/float64(len(slice))   //we need to type convert here
}

func main(){
	fmt.Println(sliceavg([]int{1,2,3,4,5,6}))
}