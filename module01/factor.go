package module01

// Factor takes in a list of primes and a number and factors that number with
// the provided primes.
//
// The returned numbers can be in any order; tests will sort them in increasing
// order to make checking easier.
//
// Bonus: Any remainder (aside from 1) that can't be factored will be treated as
// a prime in the results.
//
// Examples:
//
//	Factor([]int{2,3,5}, 30) // []int{2,3,5}
//	Factor([]int{2,3,5}, 28) // []int{2,2,7}
//	Factor([]int{2,3,5}, 720) // []int{2,2,2,2,3,3,5}
//
// Examples with remainders:
//
//	Factor([2,5], 30) // []int{2,5,3}
//	Factor([3,5], 720) // []int{3,3,5,16}
//	Factor([], 4) // []int{4}
func Factor(primes []int, number int) []int {
	out := make([]int, 0)
	for _, p := range primes {
		for number%p == 0 {
			number /= p
			out = append(out, p)
		}
	}
	return out
	/*
			Prime Factorization by Trial Division
		Say you want to find the prime factors of 100 using trial division. Start by testing each integer to see if and how often it divides 100 and the subsequent quotients evenly. The resulting set of factors will be prime since, for example, when 2 is exhausted all multiples of 2 are also exhausted.

		Find the prime factors of 100:
		100 ÷ 2 = 50; save 2
		50 ÷ 2 = 25; save 2
		25 ÷ 2 = 12.5, not evenly so divide by next highest number, 3
		25 ÷ 3 = 8.333, not evenly so divide by next highest number, 4
		But, 4 is a multiple of 2 so it has already been checked, so divide by next highest number, 5
		25 ÷ 5 = 5; save 5
		5 ÷ 5 = 1; save 5
		List the resulting prime factors as a sequence of multiples, 2 x 2 x 5 x 5 or as factors with exponents, 22 x 52.

	*/

	return nil
}
