package example1

type Shape struct {
	DrawAPI DrawAPI
}

func (s *Shape) Draw() {

}

func (s *Shape) New(drawApi DrawAPI) {
	s.DrawAPI = drawApi
}
