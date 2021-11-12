package example1

import "fmt"

type Cricket struct {
	Game
}

func (g *Cricket) Initialize() {
	fmt.Println("cricket init")
}
func (g *Cricket) startPlay() {
	fmt.Println("cricket start")
}
func (g *Cricket) endPlay() {
	fmt.Println("cricket end")
}

func (g *Cricket) Play() {
	g.Game.Play()
}
