package main
import "fmt"

func main(){
	words:=[]string{"raja","kaja","maja","apple","apple","kaja","raja","maja","Teja","TEJA"}

	m:=map[string]int{}
	for _,val:= range words{
		m[val]++
	}
	fmt.Println(m)
}