package main
import "fmt"

func main(){
	categories:=[]string{"cat", "dog", "cat", "mouse", "dog","giraffe","lion"}
	m:=map[string]int{}
	count:=0
	for _,val:=range categories{
		if _,exists:=m[val];!exists{
			m[val]=count
			count++
		}
	}
	fmt.Println("Final Label encoding: ",m)
}