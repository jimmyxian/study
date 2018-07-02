package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func bubbleSort(array []int) {
	l := len(array)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if array[j] > array[j+1] {
				array[j+1], array[j] = array[j], array[j+1]
			}
		}
	}
}

func selectionSort(array []int) {
	n := len(array)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if array[min] > array[j] {
				min = j
			}
		}
		if min != i {
			array[i], array[min] = array[min], array[i]
		}
	}
}

func insertionSort(array []int) {
	n := len(array)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
}

func partion(array []int) int {
	mid := array[0]
	head := 0
	tail := len(array) - 1
	i := 1
	for head < tail {
		if array[i] > mid {
			array[i], array[tail] = array[tail], array[i]
			tail--
		} else {
			array[i], array[head] = array[head], array[i]
			head++
			i++
		}
	}
	array[head] = mid
	return head
}

func quickSort(array []int) {
	if len(array) <= 1 {
		return
	}
	index := partion(array)
	quickSort(array[:index])
	quickSort(array[index+1:])
}

func quickSort2(array []int) {
	st := stack.New()

	st.Push(0)
	st.Push(len(array) - 1)

	for st.Len() != 0 {
		right := st.Pop().(int)
		left := st.Pop().(int)
		if left < right {
			index := partion(array[left : right+1])
			index += left
			st.Push(left)
			st.Push(index)
			st.Push(index + 1)
			st.Push(right)
		}
	}
}

func Merge(array []int, mid int) {
	n := len(array)
	temp := make([]int, n)
	i := 0
	j := mid + 1
	index := 0
	for i <= mid && j <= n-1 {
		if array[i] < array[j] {
			temp[index] = array[i]
			i++
		} else {
			temp[index] = array[j]
			j++
		}
		index++
	}

	for i <= mid {
		temp[index] = array[i]
		index++
		i++
	}

	for j <= n-1 {
		temp[index] = array[j]
		index++
		j++
	}

	for k, v := range temp {
		array[k] = v
	}
	fmt.Println(temp)
}

func MergeSort(array []int) {
	if len(array) == 1 {
		return
	}

	mid := len(array) / 2
	MergeSort(array[:mid])
	MergeSort(array[mid:])
	Merge(array, mid-1)
}

func main() {
	data := []int{3, 1, 5, 7, 9, 2, 8}
	fmt.Println("bubbleSort:")
	bubbleSort(data)
	fmt.Println(data)

	data = []int{3, 1, 5, 7, 9, 2, 8}
	fmt.Println("selectionSort:")
	selectionSort(data)
	fmt.Println(data)

	data = []int{3, 1, 5, 7, 9, 2, 8}
	fmt.Println("insertionSort:")
	insertionSort(data)
	fmt.Println(data)

	data = []int{3, 1, 5, 7, 9, 2, 8, 6, 4}
	fmt.Println("quickSort:")
	quickSort(data)
	fmt.Println(data)

	data = []int{3, 1, 5, 7, 9, 2, 8, 6, 4}
	fmt.Println("quickSort2:")
	quickSort2(data)
	fmt.Println(data)

	data = []int{3, 1, 5, 7, 9, 2, 8, 6, 4}
	fmt.Println("mergeSort:")
	MergeSort(data)
	fmt.Println(data)
}
