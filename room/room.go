package room

import (
	"main/room/layout"
	"math/rand"
	"time"
)

type Room struct {
	Layout layout.RoomLayout
}

func GenerateRoom() Room {
	aRoom := Room{}

	rand.Seed(time.Now().UnixNano())
	roomTypeIndex := rand.Intn(len(layout.Layouts))

	switch layout.Layouts[roomTypeIndex] {
	case layout.Circle:
		aRoom.Layout = layout.GenerateCircle()
	case layout.Rectangle:
		aRoom.Layout = layout.GenerateRect()
	}

	return aRoom
}
