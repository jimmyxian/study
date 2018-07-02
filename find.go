package main

import (
	"fmt"
)

func findWorld(find string, msg string) int {
	bFind := []byte(find)
	bMsg := []byte(msg)
	foundIndex := -1
	for i := 0; i < len(bMsg)-len(bFind); i++ {
		j := 0
		for j < len(bFind) && bMsg[i+j] == bFind[j] {
			j++
		}
		if j == len(bFind) {
			foundIndex = i
			return foundIndex
		}
	}
	return -1
}

func findWorld2(find string, msg string) int {
	bFind := []byte(find)
	bMsg := []byte(msg)

	i, j := 0, 0
	for i < len(bMsg) && j < len(bFind) {
		if bMsg[i] == bFind[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}

	if j == len(bFind) {
		return i - j
	} else {
		return -1
	}
}

func getNext(ps []byte) []int {
	next := make([]int, len(ps))
	next[0] = -1
	j, k := 0, -1
	for j < len(ps)-1 {
		if k == -1 || ps[k] == ps[j] {
			j++
			k++
			next[j] = k
		} else {
			k = next[k]
		}
	}
	return next
}

func findWorld3(find string, msg string) int {
	bFind := []byte(find)
	bMsg := []byte(msg)

	i, j := 0, 0

	next := getNext(bFind)
	for i < len(bMsg) && j < len(bFind) {
		if j == -1 || bMsg[i] == bFind[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == len(bFind) {
		return i - j
	} else {
		return -1
	}
}

func main() {
	msg := "hello jimmy xian asdffdasdfdasdf"
	find := "sdff"
	index := findWorld3(find, msg)
	bMsg := []byte(msg)

	fmt.Println(string(bMsg[index : index+len(find)]))
}
