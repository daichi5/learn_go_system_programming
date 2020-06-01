package main

import (
	crand "crypto/rand"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"math/rand"
)

func main() {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	src := rand.NewSource(seed.Int64())
	rng := rand.New(src)
	a := make([]byte, 20)
	rng.Read(a)
	fmt.Println(hex.EncodeToString(a))
}
