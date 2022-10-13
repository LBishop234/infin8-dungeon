package room

import (
	"fmt"
	"main/config"
	"math/rand"
	"time"
)

type Room struct {
	Width  int
	Height int
}

func GenerateRoom() Room {
	roomGenConfig := config.GetRoomGen()

	aRoom := Room{}

	rand.Seed(time.Now().UnixNano())

	aRoom.Width = rand.Intn(roomGenConfig.MaxWidth-roomGenConfig.MinWidth+1) + roomGenConfig.MinWidth
	aRoom.Height = rand.Intn(roomGenConfig.MaxHeight-roomGenConfig.MinHeight+1) + roomGenConfig.MinHeight

	fmt.Println(aRoom)

	return aRoom
}
