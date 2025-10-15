package main
import "fmt"

func main(){
	//var m map[string]int    //This creates a nil map (no memory yet).You can’t add values until you initialize it.
    
	//Initializing Map

	//m:=make(map[string]int)    //Using make

	/* Using map literal 
	m2:=map[string]int{
		"apple":5,
		"banana":10,
	}
	fmt.Println(m2)     */


	/* Adding and Updating elements
	m:=make(map[string]int)
	m["apple"]=5  //adding an element to the map
	m["apple"]=8  // updating the added element  */

	/* Retrieving values
	val:=m["apple"]
	fmt.Println(val)     */

	/* check if key exists
	Because if a key doesn’t exist, you’ll get the zero value (0 for ints, "" for strings, etc.).

	val,exists:=m["banana"]    //This is idiomatic go. Refer notes for clarification.
	if exists{
		fmt.Println("Banana count: ",val)
	}else{
		fmt.Println("Banana not found")
	}


	/* Deleting keys

	delete(m,"apple")
	fmt.Println(m)   */
    

	/* Iterating Over a Map
	Map iteration order is random — don’t expect sorting.

	m:=map[string]int{"apple":4,"Teja": 25,"Raja": 30}
	for key,value:=range m{
		fmt.Println(key,value)
	}        */

	/* length of the map
	m:=map[string]int{"teja:": 25, "Raja: ": 30,"kaja: ": 20}
	t:=len(m)
	fmt.Println(t)    */
}