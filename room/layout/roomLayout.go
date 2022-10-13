package layout

type RoomLayout interface {
	GetType() LayoutType
}

type LayoutType string

const (
	Circle    LayoutType = "Circle"
	Rectangle LayoutType = "Rectangle"
)

var Layouts []LayoutType = []LayoutType{Circle, Rectangle}
