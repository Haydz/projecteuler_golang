/*

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

package main

import "fmt"

func main() {

	var threeMultiples []int
	var fiveMultiples []int

	for i := 1; i < 1000; i++ {
		if (i * 3) < 1000 {
			threeMultiples = append(threeMultiples, i*3)
		}
		if (i * 5) < 1000 {
			fiveMultiples = append(fiveMultiples, i*5)
		}

	}
	fmt.Println(threeMultiples)
	fmt.Println(fiveMultiples)

}
