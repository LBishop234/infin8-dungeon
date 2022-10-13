package layout

import (
	"main/config"
	"math/rand"
	"time"
)

type CircleLayout struct {
	layoutType LayoutType
	radius     int
}

func (c CircleLayout) GetType() LayoutType {
	return c.layoutType
}

func NewCircle(radius int) CircleLayout {
	aCircle := CircleLayout{
		layoutType: Circle,
		radius:     radius,
	}

	return aCircle
}

func GenerateCircle() CircleLayout {
	circleGenConfig := config.GetCircleGenConfig()

	rand.Seed(time.Now().UnixNano())
	randRadius := rand.Intn(circleGenConfig.MaxRadius-circleGenConfig.MinRadius+1) + circleGenConfig.MinRadius

	return NewCircle(randRadius)
}
