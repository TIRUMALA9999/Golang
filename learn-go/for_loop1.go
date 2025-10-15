/* Now let’s make it a bit more interesting by trying a loop 
    with a condition only — Go’s version of a while loop. */

package main
import "fmt"

func main(){
	i:=10
	for i>0{
		fmt.Println(i)
		i--
	}
}