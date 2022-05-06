package module

import (
	"sort"
)

func CheckPermutation(world string, wo string) bool {
	if len(world) != len(wo) {
		return false
	}
	c := []byte(world)
	wc := []byte(wo)

	sort.Slice(c, func(i, j int) bool {
		return c[i] < c[j]
	})

	sort.Slice(wc, func(i, j int) bool {
		return wc[i] < wc[j]
	})

	for i := 0; i < len(world); i++ {
		if c[i] != wc[i] {
			return false
		}
	}

	return true
}
