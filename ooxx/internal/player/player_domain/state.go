package player_domain

type State int

const (
	Blank State = iota
	O
	X
)

func (s *State) SetO() {
	*s = O
}

func (s *State) SetX() {
	*s = X
}

func (s *State) SetBlank() {
	*s = Blank
}
