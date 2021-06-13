package bracketizer

// SweepForBrackets - evaluates a string for appropriately paired brackets, if not paired == false, if paired == true
func SweepForBrackets(s string) bool {

	openBrackets := 0
	closeBrackets := 0

	for _, c := range s {
		// c is a rune, first need to convert back to string before checking it
		if string(c) == "{" {
			openBrackets++
		} else if string(c) == "}" {
			closeBrackets++
		}
		// if there are ever more closing brackets than opening we know the string is dirty
		if closeBrackets > openBrackets {
			return false
		}
	}
	return openBrackets == closeBrackets
}
