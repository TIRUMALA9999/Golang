package main
import "fmt"

// In this x,y are named return values
func split(sum float64) (x,y float64){
	x=sum*4/9
	y=sum-x
	return   //returned named variables x and y
}

func main(){
	fmt.Println(split(10.0))
}