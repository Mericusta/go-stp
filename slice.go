package stp

func Compare[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	l := len(s1)
	for index := 0; index != l; index++ {
		if s1[index] != s2[index] {
			return false
		}
	}
	return true
}
