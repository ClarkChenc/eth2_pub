package algorithm

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomList(arr []int, size int) []int {
	if size >= len(arr) {
		return arr
	}

	ret := make([]int, size)
	for i := 0; i < size; i++ {
		ret[i] = arr[i]
	}

	rand.Seed(time.Now().Unix())
	for i := size + 1; i < len(arr); i++ {
		weight := rand.Intn(i)
		if weight < size {
			ret[weight] = arr[i]
		}
	}

	return ret
}

func TestRandom() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	size := 5

	ret := GetRandomList(arr, size)
	fmt.Println(ret)
}
