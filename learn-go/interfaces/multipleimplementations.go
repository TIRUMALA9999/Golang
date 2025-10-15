package main
import "fmt"

type Shape interface{
	Area() float64
}

type Rectangle struct{
	width,height float64
}

type Circle struct{
	radius float64
}

func (r Rectangle) Area() float64{
	return r.width*r.height
}

func (c Circle) Area() float64{
	return c.radius*3.14*c.radius
}

func PrintArea(s Shape){
	fmt.Println("Area:",s.Area())
}

func main(){
	r:=Rectangle{width:3,height:4}
	c:=Circle{radius:5}

	PrintArea(r)
	PrintArea(c)
}
