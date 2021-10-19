package encoders

import (
	"bytes"
	"image"

	"github.com/gen2brain/x264-go"
)

type H264Encoder struct {
	buffer   *bytes.Buffer
	encoder  *x264.Encoder
	realSize image.Point
}

const h264SupportedProfile = "3.1"

func (e *H264Encoder) VideoSize() (image.Point, error) {
	return e.realSize, nil
}
