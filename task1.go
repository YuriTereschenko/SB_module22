package main

import (
	"fmt"
	"math/rand"
	"time"
)

const arrSize = 10

func createRndArr() [arrSize]int {
	rand.Seed(time.Now().UnixNano())
	var arr [arrSize]int
	for i, _ := range arr {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func countNumAfter(arr [arrSize]int, value int) int {
	index := -1
	for i, val := range arr {
		if val == value {
			index = i
			break
		}
	}
	if index != -1 {
		return arrSize - index - 1
	}
	return 0
}

func main() {
	rndArr := createRndArr()
	fmt.Println(rndArr)
	desiredNum := rndArr[arrSize/2]
	fmt.Printf("There are %v numbers after %v", countNumAfter(rndArr, desiredNum), desiredNum)
}
