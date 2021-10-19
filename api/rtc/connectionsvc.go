package rtc

import (
	"fmt"

	"github.com/krishpranav/webstream/api/display"
	"github.com/krishpranav/webstream/api/encoders"
)

type RemoteScreenService struct {
	stunServer      string
	videoService    display.Service
	encodingService encoders.Service
}

func NewRemoteScreenService(stun string, video display.Service, enc encoders.Service) Service {
	return &RemoteScreenService{
		stunServer:      stun,
		videoService:    video,
		encodingService: enc,
	}
}

func hasElement(haystack []string, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func (svc *RemoteScreenService) CreateRemoteScreenConnection(screenIx int, fps int) (RemoteScreenConnection, error) {
	screens, err := svc.videoService.Screens()
	if err != nil {
		return nil, err
	}

	if screenIx < 0 || screenIx > len(screens) {
		screenIx = 0
	}
	screen := screens[screenIx]
	screenGrabber, err := svc.videoService.CreateScreenGrabber(screen, fps)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if len(screens) == 0 {
		return nil, fmt.Errorf("No available screens")
	}

	rtcPeer := newRemoteScreenPeerConn(svc.stunServer, screenGrabber, svc.encodingService)
	return rtcPeer, nil
}
