package main
import "fmt"

func append_to_slice(slice []string) []string {
	fmt.Println("Before Slice: ", slice, "len:", len(slice), "cap:", cap(slice))
	slice = append(slice, "grape", "mango", "Orange")
	fmt.Println("After append: ", slice, "len:", len(slice), "cap:", cap(slice))
	return slice
}

func main() {
	s := []string{"Teja"}
	s = append_to_slice(s) // assign the returned slice back
	fmt.Println("In main: ", s, "len:", len(s), "cap:", cap(s))
}
