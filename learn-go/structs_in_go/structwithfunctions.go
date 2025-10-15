package main

import "fmt"

/*type person struct{
	name string
	age int
}

func printperson(p person){
	fmt.Println(p.name,p.age)
}

func main(){
	t:=person{"Teja",25}
	printperson(t)
}   */


//Anonymous structs
//One can create structs without naming them

func main(){
	per:=struct{
		Name string
		Age int
	}{"Bob",30}
	fmt.Println(per.Name,per.Age)
}