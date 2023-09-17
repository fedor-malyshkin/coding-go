package px

// Для строки str найти самую длинную подстроку без повторяющихся символов
func longestSubstring(str string) string {
	chars := [128]int{}
	var (
		left     = 0
		right    = 0
		resLen   = 0
		resStart = 0
	)

	for right < len(str) {
		r := str[right]
		chars[r] += 1

		for chars[r] > 1 {
			l := str[left]
			chars[l]--
			left++
		}

		// resLen = max(resLen, right
		if resLen < right-left+1 {
			resStart = left
			resLen = right - left + 1
		}
		right++
	}
	return str[resStart : resStart+resLen]
}

/*func main() {
	result1 := longestSubstring("abcabcbb")
	checkResult(result1, "abc")
}

func checkResult(actual string, expected string) {
	if actual != expected {
		fmt.Println("'" + actual + "' != '" + expected + "': Error!")
	} else {
		fmt.Println("'" + actual + "': Ok")
	}
}
*/
