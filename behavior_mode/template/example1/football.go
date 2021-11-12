package example1

import "fmt"

type Football struct {
	Game
}

func (g *Football) Initialize() {
	fmt.Println("football init")
}
func (g *Football) startPlay() {
	fmt.Println("football start")
}
func (g *Football) endPlay() {
	fmt.Println("football end")
}
func (g *Football) Play() {
	g.Game.Play()
}
