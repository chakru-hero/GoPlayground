package main
import (
	basicFuncs "Playground1/basicFuncs"
	"fmt"
)
func main(){
	var x int16;
	fmt.Println("1. Swap two Numbers.")
	fmt.Println("2. Find longest non repeating substring.")
	fmt.Print("Select an Option : ");
	 fmt.Scan(&x); 
	switch x{
		case 1  : 
			basicFuncs.SwapTwoNumbers();
		case 2 : 
			basicFuncs.NonRepeatingSubString();
		default : 
			fmt.Println("Select an option from above");
	}
}
