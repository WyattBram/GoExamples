package itteration

func RepeatNTimes(chr string, n int) string {
	ret := ""

	for i := 0; i < n; i++ {
		ret += chr
	}
	return ret

}
