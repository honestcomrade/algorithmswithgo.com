package module01

import (
	"fmt"
	"strconv"
	"strings"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {

	s := make([]string, 0)
	for i := 1; i <= n; i++ {

		if i%3 == 0 && i%5 == 0 {
			s = append(s, "Fizz Buzz")
		} else if i%3 == 0 {
			s = append(s, "Fizz")
		} else if i%5 == 0 {
			s = append(s, "Buzz")
		} else {
			str := strconv.Itoa(i)

			s = append(s, str)
		}
	}
	fmt.Println(strings.Join(s, ", "))
}
