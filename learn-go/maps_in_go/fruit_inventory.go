package main
import "fmt"

func main(){
	m:=map[string]int{"apple":5,"berries":10,"bananas": 43,"oranges": 33,"grapes":100}
	m["mango"] = 2
    m["carrot"] = 34
    m["beans"] = 56
    m["roses"] = 78
    m["plums"] = 76
	m["apple"]=134
	delete(m,"berries")
	fmt.Println("map lenght: ",len(m))
	fmt.Println("map Elements: ",m)
}