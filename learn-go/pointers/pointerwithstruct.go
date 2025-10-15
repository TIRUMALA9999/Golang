package main
import "fmt"

type Person struct{
	Name string
	Age int
}

func UpdateAge(p *Person,newAge int){
	p.Age=newAge
}

func main(){
	person:=Person{Name: "Alice",Age:20}
	fmt.Println("Before:",person.Age)
	UpdateAge(&person,30)
	fmt.Println("After:",person.Age)
}

