package randomhelper

import "math/rand"

func RandomDistribute(keys []string, total int) map[string]int {
	ret := make(map[string]int, len(keys))
	l := len(keys)
	for i := 0; i < l-1; i++ {
		curr := rand.Int31n(int32(total))
		ret[keys[i]] = int(curr)
		total = total - int(curr)
	}
	ret[keys[l-1]] = total
	return ret
}

// @links https://yourbasic.org/golang/shuffle-slice-array/
