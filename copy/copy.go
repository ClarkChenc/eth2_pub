package copy

import "fmt"

type Data struct {
	data1 []int
	data2 []int
	num   int
}

func CopyTest() {
	org := Data{}
	org.data1 = []int{1, 2, 3, 4}
	org.num = 3

	copy := org
	copy.data1[0] = 2
	copy.num = 4
	copy.data2 = []int{2, 4, 6, 8}

	fmt.Println("org: ", org)
	fmt.Println("copy: ", copy)
}
