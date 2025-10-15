package main
import "fmt"

type DataPoint struct{
	Features []float64
	Label string
}


func main(){
	dataset:=[]DataPoint{
		{Features: []float64{5.1, 3.5, 1.4, 0.2}, Label: "Setosa"},
		{Features: []float64{7.0, 3.2, 4.7, 1.4}, Label: "Versicolor"},
		{Features: []float64{6.3, 3.3, 6.0, 2.5}, Label: "Virginica"},
		{Features: []float64{5.0, 3.6, 1.4, 0.2}, Label: "Setosa"},
	}

	fmt.Println("Dataset:")
	for _,dp:=range dataset{
		fmt.Println("Features: ",dp.Features,"Label: ",dp.Label)
	}

	labelCount:=map[string]int{}

	for _,dp:=range dataset{
		labelCount[dp.Label]++
	}

	fmt.Println("\n Label Counts:")
	for label,count:=range labelCount{
		fmt.Printf("%s -> %d\n", label, count)
	}
}