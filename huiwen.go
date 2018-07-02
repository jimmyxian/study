package main

import (
	"fmt"
)

func findLongest(msg []byte) {
	maxlen := 0
	start := 0

	for i := 0; i < len(msg); i++ {
		for j := i + 1; j < len(msg); j++ {
			k := i
			l := j
			for k < l {
				if msg[k] != msg[l] {
					break
				}
				k++
				l--
			}
			if k >= l && j-i > maxlen {
				maxlen = j - i + 1
				start = i
			}
		}
	}
	if maxlen > 0 {
		fmt.Println(string(msg[start : start+maxlen]))
	}
}

func main() {
	findLongest([]byte("asdfbbcbbasd"))
}
