package operator

import "fmt"

const (
	sigSize = 127
)

var (
	sigs [(sigSize + 31) / 32]uint32
)

func set(sig int) {
	sigs[sig/32] |= 1 << (sig & 31)
}

func want(sig int) bool {
	return sigs[sig/32]&(1<<(sig&31)) != 0
}

func clear(sig int) {
	sigs[sig/32] &= ^(1 << (sig & 31))
}

func TestSigOperator() {
	// set 0x0001
	set(1)
	fmt.Printf("set (1): %08b\n", sigs)

	clear(0)
	fmt.Printf("clear(0): %08b\n", sigs)

	set(2)
	fmt.Printf("set (2): %08b\n", sigs)

	clear(1)
	fmt.Printf("clear(1): %08b\n", sigs)

	fmt.Printf("want(0): %v, %08b\n", want(0), sigs)

	fmt.Printf("want(1): %v, %08b\n", want(1), sigs)

	fmt.Printf("want(2): %v, %08b\n", want(2), sigs)
}
