package  main 

func romanToInt(s string) int {

    val := 0

	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := 0; i < len(s); i++ {
		curr := roman[string(s[i])]
		if i+1 < len(s) {
			next := roman[string(s[i+1])]
			if curr < next {
				val -= curr // Subtract the current value
			} else {
				val += curr // Add the current value
			}
		} else {
			val += curr // Add the last value
		}
	}

	return val
    
}