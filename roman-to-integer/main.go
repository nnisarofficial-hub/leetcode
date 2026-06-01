package main

func romanToInt(s string) int {
	romans := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	chars := []rune(s)
	sum := 0
	length := len(chars)

	for index := 0; index < length; index++ {
		currentKey := (chars[index])
		currentVal := romans[currentKey]

		if index+1 < length {
			nextKey := (chars[index+1])
			nextVal := romans[nextKey]

			if currentVal < nextVal {
				sum -= currentVal
				continue
			}
		}
		sum += currentVal
	}
	return sum
}
