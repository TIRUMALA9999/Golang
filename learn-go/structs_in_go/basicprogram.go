package main
import "fmt"

type Book struct{
	Title string
	Author string
}


func main(){
	b:=Book{Title: "Law of Attraction",Author:"Teja"}
	fmt.Println(b)
}