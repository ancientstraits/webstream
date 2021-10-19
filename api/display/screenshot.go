package display

//import "image"
//import "github.com/kbinani/screenshot"

import (
	"image"
	"github.com/kbinani/screenshot"
	"time"
)

type VideoProvider struct{}

type ScreenGrabber struct {
	fps int
	screen Screen
	frames chan *image.RGBA
	stop chan struct{}
}

func (x *VideoProvider) Screens() ([]Screen, error) {
	numScreens := screenshot.NumActiveDisplays()
	screens := make([]Screen, numScreens)
	for i := 0; i < numScreens; i++ {
		screens[i] = Screen{
			Index: i,
			Bounds: screenshot.GetDisplayBounds(i),
		}
	}

	return screens, nil
}

func (g *ScreenGrabber) Stop() {
	close(g.stop)
}

func (g *ScreenGrabber) Fps() int {
	return g.fps
}

func NewVideoProvider() (Service, error) {
	return &VideoProvider{}, nil
}