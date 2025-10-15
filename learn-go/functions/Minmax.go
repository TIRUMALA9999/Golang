package main
import "fmt"

func minMax(a,b int) (int,int){
	if a>b{
		return a,b
	}
	return b,a

}

func main(){
	x,y:=minMax(5345436,3466754)
	if x>y{
		fmt.Printf("x=%d is larger & y=%d is smaller\n",x,y)
	}else{
		fmt.Printf("y=%%d is larger & x=%%d is smaller",y,x)
	}
}