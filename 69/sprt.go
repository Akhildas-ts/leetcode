package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {

	sq := math.Sqrt(float64(x))


	return int(sq)

}

func main() {

fmt.Println(mySqrt(8))

}
