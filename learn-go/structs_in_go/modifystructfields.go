package main
import "fmt"

type car struct{
	Brand string
	Year int
}

func main(){
	c:=car{Brand:"Toyota",Year:2025}
	fmt.Println("Old value:",c.Year)
	c.Year=2026
	fmt.Println("New value:",c.Year)
}