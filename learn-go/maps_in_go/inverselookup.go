package main
import "fmt"

func popvalue(a int) []string{
	cities := map[string]int{
    "New York": 8419000,
    "Los Angeles": 3980000,
    "Chicago": 2716000,
    "Houston": 2328000,
    "Phoenix": 1690000,
    "San Diego": 3980000,
	}
	t:=[]string{}
	for key,val:=range cities{
		if val==a{
			t=append(t,key)
		}
	}
	return t
}

func main(){
	fmt.Println(popvalue(3980000))
}