package main
import "fmt"

type Student struct{
	Name string
	Grade string
}

func (s Student) Display(){
	fmt.Println(s)
}

func (s *Student) UpdateGrade(newGrade string){
	s.Grade=newGrade
}

func main(){
	s1:=Student{"Teja","A"}
	s2:=Student{"Raja","A++"}
	s3:=Student{"Kaja","A+"}

	s1.Grade="B"

	s1.Display()
	s2.Display()
	s3.Display()
}