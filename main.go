package main

import (
	"fmt"

	poseidongold "github.com/okx/poseidongold/go"
)

func main() {
	var input [8]uint64 = [8]uint64{0, 0, 0, 0, 0, 0, 0, 0}
	var capacity [4]uint64 = [4]uint64{0, 0, 0, 0}
	var result [4]uint64

	poseidongold.HashWithResult(&input, &capacity, &result)
	fmt.Println(result)
}
