package reverse_string

/* https://leetcode.com/problems/reverse-string/ */

func solution(s []byte) {
	start := 0
	end := len(s) - 1

	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
