package main
import (
	"fmt"
)
func main() {
	x := 5
	y := 7
	sum := x + y
	if sum > 8{ 
	fmt.Println("Greater than 8")
	} else if sum < 2 {
	fmt.Println("less than 2")
	} else { 
	fmt.Println("wrong choice")
	}
}