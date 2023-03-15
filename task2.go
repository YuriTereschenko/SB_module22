package main

import "fmt"

const arrSize2 = 12

func find(arr [arrSize2]int, value int) int {
	minI := 0
	maxI := arrSize2 - 1
	for maxI > minI {
		curIndex := (maxI + minI) / 2
		if arr[curIndex] == value {
			for {
				if arr[curIndex-1] == value {
					curIndex--
				} else {
					return curIndex
				}
			}

		} else if arr[curIndex] > value {
			maxI = curIndex - 1
		} else {
			minI = curIndex + 1
		}
	}
	return -1
}

func main() {
	arr := [arrSize2]int{1, 1, 1, 2, 2, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(find(arr, 2))
}
