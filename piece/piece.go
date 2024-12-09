package piece


type Turn int

const (
	White Turn = iota
	Black
)

type Piece struct {
	Name    string
	Color   Turn
	Display string
	HasMoved bool
}


