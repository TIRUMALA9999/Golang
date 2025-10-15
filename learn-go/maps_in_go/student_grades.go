package main

import "fmt"



func main(){

    scores:=[]([]int){{67},{78},{76},{98},{87},{69}}

    names:=[]string{"Raja","Teja","Kaja","maja","baja","vaja"}

    m:=map[string]([]int){}

    for i,name:=range names{

        m[name]=[]int(scores[i])       

    }

    fmt.Println(m)

    m["Raja"]=[]int{76,65,82}

    m["Teja"]=[]int{79,88,99}

    for key,val:= range m{
		sum:=0
		for _,scores:=range val{
			sum+=scores
		}
		average:=float64(sum)/float64(len(val))
		fmt.Println(key,":",average)        

    }   

}