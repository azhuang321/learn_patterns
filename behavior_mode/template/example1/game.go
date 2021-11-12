package example1

type Gamer interface {
	Initialize()
	startPlay()
	endPlay()
}

type Game struct{}

func (g *Game) Initialize() {}
func (g *Game) startPlay()  {}
func (g *Game) endPlay()    {}
func (g *Game) Play() {
	g.Initialize()
	g.startPlay()
	g.endPlay()
}

func Play(gamer Gamer) {
	gamer.Initialize()
	gamer.startPlay()
	gamer.endPlay()
}
