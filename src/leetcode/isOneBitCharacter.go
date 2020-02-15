package leetcode

// IsOneBitCharacter link https://leetcode-cn.com/problems/1-bit-and-2-bit-characters/
func IsOneBitCharacter(bits []int) bool {
	for i := 0; i < len(bits)-1; {
		if bits[i] == 0 {
			i++
		} else {
			if i == len(bits)-2 {
				return false
			}
			i += 2
		}
	}
	return true
}
