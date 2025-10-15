package main
import "fmt"

type Address struct{
	City string
	State string
}

type Person struct{
	Name string
	Age int
	Address Address          //we are passing address struct that is displayed above
}

/*func main(){
	p := Person{"Teja", 21, Address{"Vijayawada", "AP"}}
	fmt.Println(p.Address.City) // Vijayawada
} */

func main(){
	//structs & Slices/Maps
	people:=[]Person{
		{"Teja", 21, Address{"Vijayawada", "AP"}},
		{"Alice", 25, Address{"Hyderabad", "TS"}},
	}

	for _,person:=range people{
		fmt.Println(person.Name, person.Address.City)
	}
}