package main
import "fmt"

func main(){
	m:=map[string]string{"Teja":"yegineni@gmail.com","raja":"raja@hotmail.com","kaja":"kaja@outlook.com","jaja":"","baja":""}
	for name,email:=range m{
		if email==""{
			delete(m,name)
		}
	}
	fmt.Println(m)

}