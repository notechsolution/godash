package lo

// Check whether the input string is number or not. using the ascii Dec char value to check.
func IsNumber(number string) bool {
	for _, c := range number {
		if c < 48 || c > 57 {
			return false;
		}
	}
	return true;
}
