package main
import "fmt"

type Address struct{
	City string
	ZipCode string
}

type Person struct{
	Name string
	Age int
	Address Address
}

func (p *Person) UpdateCity(NewCity string){
	p.Address.City=NewCity
}

func main(){
	p1:=Person{"Teja",25,Address{"Hyderabad","75070"}}
	p2:=Person{"Raja",30,Address{"Vijayawada","524224"}}
	fmt.Printf("Name: %s, Age: %d, City: %s, Zip: %s\n", p1.Name, p1.Age, p1.Address.City, p1.Address.ZipCode)
	fmt.Printf("Name: %s, Age: %d, City: %s, Zip: %s\n", p2.Name, p2.Age, p2.Address.City, p2.Address.ZipCode)

	// Update city for p1
	p1.UpdateCity("Chicago")

	// Print updated data
	fmt.Printf("Updated -> Name: %s, Age: %d, City: %s, Zip: %s\n", p1.Name, p1.Age, p1.Address.City, p1.Address.ZipCode)
}
