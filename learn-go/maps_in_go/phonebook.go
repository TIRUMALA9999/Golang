package main
import "fmt"

func main(){
	m:=map[string]int{"Teja": 9409770152,"Srinu": 9441794486,"Nag mama":9490273784,"Raja":9547389028}
	name:="Srinu"
	if number,exists:=m[name];exists{
		fmt.Println("Name exists! Phone number is: ",number)
	}else{
		fmt.Println("Name doesn't exist")
	}
	fmt.Println()


	name="babu"
	if number,exists:=m[name];exists{
		fmt.Println("Name exists!",number)
	}else{
		fmt.Println("Name doesn't exist")
	}

	fmt.Println()

	fmt.Println("Complete phonebook:")
	for key,val:=range m{
		fmt.Println(key,":",val)
	}
}