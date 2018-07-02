package main

import (
	"container/list"
	"fmt"
)

func QuanList(msg []byte) {
	Quan(msg, 0)
}

func IsSwap(msg []byte, p byte) bool {
	for _, v := range msg {
		if v == p {
			return false
		}
	}
	return true
}

func Quan(msg []byte, index int) {
	if index == len(msg)-1 {
		fmt.Println(string(msg))
	} else {
		for i := index; i < len(msg); i++ {
			if IsSwap(msg[i+1:], msg[i]) == true {
				msg[i], msg[index] = msg[index], msg[i]
				Quan(msg, index+1)
				msg[i], msg[index] = msg[index], msg[i]
			}
		}
	}
}

func Comb2(msg []byte, number int, rst *list.List) {
	if number == 0 {
		for e := rst.Front(); e != nil; e = e.Next() {
			fmt.Print(string(e.Value.(byte)))
		}
		fmt.Println("----")
		return
	}
	if len(msg) == 0 {
		return
	}

	rst.PushBack(msg[0])
	Comb2(msg[1:], number-1, rst)
	rst.Remove(rst.Back())
	Comb2(msg[1:], number, rst)
}

func Comb(msg []byte) {
	rst := list.New()
	for i := 1; i <= len(msg); i++ {
		Comb2(msg, i, rst)
	}
}

func main() {
	//QuanList([]byte("abb"))

	Comb([]byte("abc"))
}
