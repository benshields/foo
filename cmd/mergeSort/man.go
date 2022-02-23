package main

import "fmt"

func main() {
	data := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	fmt.Printf("%v\n%v\n", data, MergeSortSeq(data))
	fmt.Printf("%v\n", MergeSort(data))
	fmt.Printf("%v\n", MergeSort2(data))
}

////////////////////////////////////////////////////////

func MergeSort(data []int) []int {
	merged := make(chan []int)
	go mergeSort(data, merged)
	return <-merged
}

func mergeSort(data []int, out chan []int) {
	if len(data) <= 1 {
		out <- data
		return
	}

	mid := len(data) / 2

	left := make(chan []int)
	go mergeSort(data[:mid], left)

	right := make(chan []int)
	go mergeSort(data[mid:], right)

	out <- merge(<-left, <-right)
}

func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	return merged
}

////////////////////////////////////////////////////////

func MergeSortSeq(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	left := MergeSortSeq(data[:mid])
	right := MergeSortSeq(data[mid:])
	return MergeSeq(left, right)
}

func MergeSeq(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	return merged
}

////////////////////////////////////////////////////////

func MergeSort2(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	done := make(chan struct{})

	var left, right []int
	go func() {
		left = MergeSort2(data[:mid])
		done <- struct{}{}
	}()

	right = MergeSort2(data[mid:])

	<-done
	return Merge2(left, right)
}

func Merge2(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	return merged
}
