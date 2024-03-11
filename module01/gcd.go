package module01

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	a, b = b, a%b
	return GCD(a, b)
	// primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}
	// f1 := dedupe(Factor(primes, a))
	// f2 := dedupe(Factor(primes, b))
	// out := findSharedElements(f1, f2)
	// // out = append(out)
	// if len(out) == 0 {
	// 	return 1
	// }
	// return sum(out)
}

func dedupe(a []int) []int {
	occurences := map[int]bool{}
	for _, s := range a {
		occurences[s] = true
	}
	keys := make([]int, len(occurences))
	i := 0
	// do this with a while?
	for k := range occurences {
		keys[i] = k
		i++
	}
	return keys
}

// TODO fix this function
// needs to choose only one copy of any specific value from each array once
// so dedupe
func findSharedElements(slice1, slice2 []int) []int {
	sharedElements := []int{}
	elementMap := map[int]bool{}

	for _, element := range slice1 {
		elementMap[element] = true
	}

	for _, element := range slice2 {
		if _, ok := elementMap[element]; ok {
			sharedElements = append(sharedElements, element)
		}
	}

	return sharedElements
}

func sum(slice []int) int {
	sum := 0
	for _, s := range slice {
		sum += s
	}
	return sum
}
