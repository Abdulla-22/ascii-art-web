package Web_Art

func Line_num(num byte) int {
	// Subtract 32 from the giving value to start from the top of the list because the value of the first character is 32.
	// Then multiply by 9 because the character has 8 rose plus the new line "\n".
	// Then plus 2 to remove the extra plank place in the top.
	return (9)*((int(num)) - 32) + 2
}
