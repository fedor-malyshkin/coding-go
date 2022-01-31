package p844

/*
https://leetcode.com/problems/backspace-string-compare/
*/
func backspaceCompare(s string, t string) bool {
	i := len(s) - 1
	j := len(t) - 1
	var (
		skipS = 0
		skipT = 0
	)

	for i >= 0 || j >= 0 { // While there may be chars in build(S) or build (T)
	loopI:
		for i >= 0 { // Find position of next possible char in build(S)
			switch {
			case s[i] == '#':
				skipS++
				i--
			case skipS > 0:
				{
					skipS--
					i--
				}
			default:
				break loopI
			}
		}
	loopJ:
		for j >= 0 { // Find position of next possible char in build(T)
			switch {
			case t[j] == '#':
				{
					skipT++
					j--
				}
			case skipT > 0:
				{
					skipT--
					j--
				}
			default:
				break loopJ
			}
		}
		// If two actual characters are different
		if i >= 0 && j >= 0 && s[i] != t[j] {
			return false
		}
		// If expecting to compare char vs nothing
		if (i >= 0) != (j >= 0) {
			return false
		}
		i--
		j--
	}
	return true
}
