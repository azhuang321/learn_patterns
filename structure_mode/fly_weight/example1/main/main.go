package main

import (
	"crypto/rand"
	"go-patterns/structure_mode/fly_weight/shape"
	"math/big"
)

func main() {
	for i := 0; i < 20; i++ {
		circle := new(shape.ShapeFactory).GetCircle(getRandColor())
		circle.SetX(getRandInt())
		circle.SetY(getRandInt())
		circle.SetRadius(getRandInt())
		circle.Draw()
	}
}

func getRandColor() string {
	colorSlice := []string{"red", "blue", "green", "yellow", "white", "black"}
	num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(colorSlice))))
	return colorSlice[num.Int64()]
}

func getRandInt() int {
	num, _ := rand.Int(rand.Reader, big.NewInt(100))
	return int(num.Int64())
}
