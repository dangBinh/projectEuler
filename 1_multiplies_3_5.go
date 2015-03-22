package main 

import "fmt"

/* If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000. */
// normal way
func solutionOne() {
	k := 0
	s := 0
	for i := 0; i < 1000; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			s = s + i
			k = k + 1
		}
	}
	return s
}
const target = 999
// efficient way 
func devisionBy(n) {
	p := target / n
	return n * (p * (p + 1)) / 2
}

func solutionTwo() {
	result := devisionBy(3) + devisionBy(5) - devisionBy(15)
	return result
}

func main() {
	result := solutionOne
	fmt.Println(result)
	result = solutionTwo
	fmt.Println(result)
}