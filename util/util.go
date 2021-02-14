package util

import (
	"log"
	"strconv"
	"strings"
)

func strs2ints(s []string) []int {
	a := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		n, err := strconv.Atoi(s[i])
		if err != nil {
			log.Fatal(err)
		}
		a[i] = n
	}

	return a
}

// GetVersions ex) "1.2.3" => 1, 2, 3
func GetVersions(s string) (int, int, int) {
	versions := strings.Split(s, ".")
	v := strs2ints(versions)
	major := v[0]
	minor := v[1]
	patch := v[2]

	return major, minor, patch
}
