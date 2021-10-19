package display

//import "image"
//import "github.com/kbinani/screenshot"

import (
	"github.com/kbinani/screenshot"
	"image"
)

type VideoProvider struct{}

func (x *VideoProvider) CreateScreenGrabber(screen Screen, fps int) (ScreenGrabber, error) {
	panic("implement me")
}

type XScreenGrabber struct {
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

func (g *XScreenGrabber) Stop() {
	close(g.stop)
}

func (g *XScreenGrabber) Fps() int {
	return g.fps
}

func NewVideoProvider() (Service, error) {
	return &VideoProvider{}, nil
}