package layout

import (
	"main/config"
	"math/rand"
	"time"
)

type RectLayout struct {
	layoutType LayoutType
	height     int
	width      int
}

func (r RectLayout) GetType() LayoutType {
	return r.layoutType
}

func NewRect(height, width int) RectLayout {
	aRect := RectLayout{
		layoutType: Rectangle,
		height:     height,
		width:      width,
	}

	return aRect
}

func GenerateRect() RectLayout {
	rectGenConfig := config.GetRectGenConfig()

	rand.Seed(time.Now().UnixNano())
	randHeight := rand.Intn(rectGenConfig.MaxHeight-rectGenConfig.MinHeight+1) + rectGenConfig.MinHeight
	randWidth := rand.Intn(rectGenConfig.MaxWidth-rectGenConfig.MinWidth+1) + rectGenConfig.MinWidth

	return NewRect(randHeight, randWidth)
}
