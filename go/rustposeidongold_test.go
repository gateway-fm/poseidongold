package poseidongold

import (
	"math/rand"
	"testing"

	"github.com/gateway-fm/vectorized-poseidon-gold/src/vectorizedposeidongold"
	goldenposeidon "github.com/iden3/go-iden3-crypto/goldenposeidon"
)

func TestHashWithResult(t *testing.T) {
	var SIZE int = 8

	var input [][8]uint64 = make([][8]uint64, SIZE)
	var capacity [4]uint64 = [4]uint64{0, 0, 0, 0}
	var result [4]uint64
	var passedTest bool

	rand.Seed(1)
	for i := 0; i < SIZE; i++ {
		for j := 0; j < 8; j++ {
			input[i][j] = rand.Uint64()
		}
	}

	for i := 0; i < SIZE; i++ {
		iden3HashResult, _ := goldenposeidon.Hash(input[i], capacity)
		HashWithResult(&input[i], &capacity, &result)

		passedTest = true
		for j := 0; j < 4; j++ {
			if result[j] != iden3HashResult[j] {
				passedTest = false
			}
		}

		if !passedTest {
			t.Errorf("Expected Hash %v, Actual Hash %v", iden3HashResult, result)
			break
		}
	}
}

func prepareInputs() ([8]uint64, [4]uint64) {
	var input [8]uint64 = [8]uint64{
		5577006791947779410,
		8674665223082153551,
		15352856648520921629,
		13260572831089785859,
		3916589616287113937,
		6334824724549167320,
		9828766684487745566,
		10667007354186551956,
	}
	var capacity [4]uint64 = [4]uint64{0, 0, 0, 0}
	return input, capacity
}

func BenchmarkHashWithResult(b *testing.B) {
	inputs, capacity := prepareInputs()
	var result [4]uint64

	for i := 0; i < b.N; i++ {
		HashWithResult(&inputs, &capacity, &result)
	}
}

func BenchmarkVectorizedPoseidongold(b *testing.B) {
	inputs, capacity := prepareInputs()
	var result [4]uint64

	for i := 0; i < b.N; i++ {
		vectorizedposeidongold.HashWithResult(&inputs, &capacity, &result)
	}
}
