/*

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

package main

import "fmt"

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.

func Contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func main() {

	var threeMultiples []int
	var fiveMultiples []int

	for i := 1; (i * 3) < 1000; i++ {
		//fmt.Println(i)
		threeMultiples = append(threeMultiples, i*3)
	}

	for i := 1; (i * 5) < 1000; i++ {
		//fmt.Println(i)
		fiveMultiples = append(fiveMultiples, i*5)
	}

	//initial way of solving
	// for i := 1; i < 1000; i++ {
	// 	// if (i * 3) < 1000 {
	// 	// 	threeMultiples = append(threeMultiples, i*3)
	// 	// }
	// 	if (i * 5) < 1000 {
	// 		fiveMultiples = append(fiveMultiples, i*5)
	// 	}
	// }

	//compare sum the difference of integers within the 2 slices
	for _, value := range threeMultiples {
		if Contains(fiveMultiples, value) == false {
			fiveMultiples = append(fiveMultiples, value)

		}
	}

	totalSum := 0

	for _, value := range fiveMultiples {
		//fmt.Println("value: ", value)
		totalSum = totalSum + value
	}
	fmt.Println(totalSum)

}
