
package main
import "fmt"

type Counter struct{
	value int
}


func (c *Counter) Increment(){
	c.value++
}

func main(){
	c:=Counter{}
	c.Increment()    //Go automatically does(&c).Increment()
	fmt.Println(c.value)
}