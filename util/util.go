package util

import (
	"log"
	"strconv"
)

func Strings2int(s []string) []int {
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
