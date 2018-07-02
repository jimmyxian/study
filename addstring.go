package main

import (
	"fmt"
)

func AddStringNumber(s1 string, s2 string) string {
	b1 := []byte(s1)
	b2 := []byte(s2)

	maxlen := 0
	if len(b1) > len(b2) {
		maxlen = len(b1)
	} else {
		maxlen = len(b2)
	}
	rst := make([]byte, maxlen+1)

	var carry byte = 0
	i := 0
	var temp1 byte
	var temp2 byte
	for i < len(b1) && i < len(b2) {
		if i >= len(b1) {
			temp1 = 0
		} else {
			temp1 = b1[i] - '0'
		}

		if i >= len(b2) {
			temp2 = 0
		} else {
			temp2 = b2[i] - '0'
		}

		sum := temp1 + temp2 + carry
		rst[i] = sum%10 + '0'
		carry = sum / 10
		i++
	}

	if carry != 0 {
		rst[i] = byte(carry + '0')
		return string(rst)
	} else {
		return string(rst[:maxlen])
	}
}

func AddTwo(s1 string, s2 string) string {
	b1 := []byte(s1)
	b2 := []byte(s2)

	maxlen := 0
	if len(b1) > len(b2) {
		maxlen = len(b1)
	} else {
		maxlen = len(b2)
	}
	rst := make([]byte, maxlen+1)

	i := len(b1) - 1
	j := len(b2) - 1
	k := maxlen
	var sum byte
	var carry byte = 0
	for i >= 0 && j >= 0 {
		sum = 0
		if i >= 0 {
			sum += (b1[i] - '0')
		}

		if j >= 0 {
			sum += (b2[j] - '0')
		}

		rst[k] = sum%10 + '0' + carry
		carry = sum / 10
		i--
		j--
		k--
	}
	if carry != 0 {
		rst[0] = carry + '0'
		return string(rst)
	} else {
		return string(rst[1:])
	}
}

func main() {
	rst := AddStringNumber("1239", "1234")
	fmt.Println(rst)
	rst = AddTwo("1239", "1234")
	fmt.Println(rst)
}
