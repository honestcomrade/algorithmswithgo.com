package module01

// BaseToBase takes in a number, the base it is currently
// in, and the base you want it to be converted to. It then
// returns a string representing the number in the new base.
//
// Eg:
//
//	BaseToBase("E", 16, 2) => "1110"
func BaseToBase(value string, base, newBase int) string {
	return ""
	/* 14 to base 2:

	Step 1: 14 % 2 = 0, so next digit is "0".
	  Our new number is currently "0"

	Step 2: 14 / 2 = 7, so we update our number to be 7
	  7 is > 0, so we go to step 1...

	Step 1: 7 % 2 = 1; next digit is "1"
	  Our new number is currently "10" (we add the new digit to the left)

	Step 2: 7 / 2 = 3 (round down always), so we update our number to be 3
	  3 > 0, so go to step 1

	Step 1: 3 % 2 = 1; next digit is "1"
	  Our new number is "110"

	Step 2: 3 / 2 = 1; Go to step 1

	Step 1: 1 % 1 = 1; next digit is "1"
	  Our new number is "1110"

	Step 2: 1 / 2 = 0; stop here!

	Final number is "1110", which is 14 in base 2! */
	// return strings.Join(s, "")
}

func Modulo(x int, z int) int {
	return x % z
}
