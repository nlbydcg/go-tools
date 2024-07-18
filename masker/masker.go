package masker

import "strings"

func Mobile(s string) string {
	return masker(s, '*', 3, 4)
}

func Email(s string) string {
	i := strings.Index(s, "@")
	n := 4
	if i-4 <= 0 {
		n = i - 4
		i = 0

	}
	return masker(s, '*', i, n)
}

func masker(s string, r byte, i, n int) string {
	if len(s) < i {
		return s
	}

	list := []rune(s)
	ll := len(list)
	for j := 0; j < n; j++ {
		if j+i < ll {
			list[j+i] = rune(r)
		}
	}
	return string(list)
}
