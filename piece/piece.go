package piece

// impor the turn struct here. turn here.


type Piece struct {
	Name    string
	Color   Color
	Display string
	HasMoved bool
}


type Color int

const (
	White Color = iota
	Black
)