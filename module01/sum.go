package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	ans := 0
	if numbers != nil {
		length := len(numbers)
		if length <= 0 {
			return ans
		}
		for i := 0; i <= length-1; i++ {
			amt := numbers[i]
			if amt == 0 {
				continue
			}
			ans += amt
		}
		return ans
	}
	return ans
}
