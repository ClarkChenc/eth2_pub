package decimal

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func TestDecimal() {
	decimal, err := decimal.NewFromString("47854622944193680932087931")
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(decimal)
}
