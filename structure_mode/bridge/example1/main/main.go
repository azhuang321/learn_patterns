package main

import "go-patterns/structure_mode/bridge"

func main() {
	redCircle := bridge.New(10, 100, 100, new(bridge.RedCircle))
	greenCircle := bridge.New(10, 100, 100, new(bridge.GreenCircle))

	redCircle.Draw()
	greenCircle.Draw()

}
