package main
import "fmt"

func main(){
	original := []int{1, 2, 3}
	copySlice := original
	copySlice[0] = 100
	fmt.Println(original)  // Output: [100 2 3]
	fmt.Println(copySlice)  // Output: [100 2 3]
 
}