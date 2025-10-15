package main
import "fmt"

func main(){
	words:=[]string{"go", "is", "fun", "go", "is", "fast"}

	m:=map[string]int{}

	for _,word:=range words{
		m[word]++
	}
	fmt.Println(m)
}