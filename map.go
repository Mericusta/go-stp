package stp

func Key[K comparable, V any](tm map[K]V) []K {
	ks, i := make([]K, len(tm)), 0
	for k := range tm {
		ks[i] = k
		i++
	}
	return ks
}
