package exam

// IsAnagram q
func IsAnagram(str1, str2 string) bool {

	for k := 0; k < len(str1); k++ {
		c1 := str1[k]

		found := false

		for i := 0; i < len(str2); i++ {
			c2 := str2[i]
			if c1 == c2 {
				found = true

				str2 = str2[:i] + str2[i+1:]
				i--

				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}
