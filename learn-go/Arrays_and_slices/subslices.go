package  main
import "fmt"

func subslice(s []int) []int{
	slice:=s[2:6]
	return slice
}

func main(){
	fmt.Println(subslice([]int{5,4,3,6,8,98,345,678,9,4}))
}