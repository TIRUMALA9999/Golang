package main
import "fmt"

func main(){
	scores:=[]([]int){{56},{66},{78},{85},{91},{95}}
	names:=[]string{"Teja","Raja","kaja","maja","gaja","laja"}
	m:=map[string]([]int){}

	for i,name:=range names{
		m[name]=scores[i]
	}
	fmt.Println(m)

	m["Raja"]=[]int{56,78,65}
	m["Teja"]=[]int{67,85,99}

	for name,scores:=range m{
		sum:=0
		for _,score:=range scores{
			sum=sum+score
		}
		average:=float64(sum)/float64(len(scores))
		fmt.Println(name,":",average)
	}
}