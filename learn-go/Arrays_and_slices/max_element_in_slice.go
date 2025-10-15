package main
import "fmt"

func max(slice []int) int{
	temp:=slice[0]
	for i,_:=range slice{
		if slice[i]>temp{
			temp=slice[i]
		}
	}
	return temp
}

func main(){
	fmt.Println(max([]int{2398573,2343545,689576,23953496437,345,468,23598}))
}