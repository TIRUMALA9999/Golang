package main
import "fmt"

func main(){
	
	//Declaring a struct
	type person struct{
		name string
		age int
	}

	//Using the declared struct and initializing the values
	p1:=person{name:"Teja",age:21}   //initializing using field names
	p2:=person{"Raja",22}

	p1.age=25   //modifying the elements
	fmt.Println(p1.age)    //p1.age is the accessing elements
	fmt.Println(p2.name)
}