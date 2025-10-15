package main
import "fmt"

func main(){
	var score int
	fmt.Print("Enter your score:")
	fmt.Scanln(&score)

	if score>=90{
		fmt.Println("Grade A")
	}else if score >=80{
		fmt.Println("Grade B")
	}else if score >=70{
		fmt.Println("Grade C")
	}else{
		fmt.Println("Grade D or Grade F")
	}
}