package main

import "fmt"

func getMonthlyPrice(tier string) int {
	// ?
	switch tier {
	case "basic":
		/* code */
		return 100 * 100
	case "premium":
		return 150 * 100
	case "enterprise":
		return 500 * 100
	default:
		/* code */
		return 0
	}
}


func main(){
	fmt.Println(getMonthlyPrice("basic"))
	fmt.Println(getMonthlyPrice("premium"))
	fmt.Println(getMonthlyPrice("enterprise"))

}