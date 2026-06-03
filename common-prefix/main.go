package main

func main() {

}

func longestCommonPrefix(strs []string) string {
	result := ""

	for position := 0; position < len(strs[0]); position++ {
		ref := strs[0][position]

		for _, word := range strs {
			if position >= len(word) || word[position] != ref {

				return result
			}
		}
		result += string(ref)
	}

	return result

}
