package main
import "fmt"

func copy(s []int, t []int) []int{
	for i,_:= range s{
		t=append(t,s[i])
	}
	return t
}

func main(){
	res:=copy([]int{4,2,5,7,545,678,3,89,39},[]int{})
	fmt.Println(res)
}