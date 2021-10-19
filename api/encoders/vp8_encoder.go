package encoders

import "C"
import (
	"bytes"
	"image"
)

const keyFrameInterval = 10

type VP8Encoder struct {
	buffer   *bytes.Buffer
	realSize image.Point
}

func newVP8Encoder(size image.Point, frameRate int) (Encoder, error) {
	buffer := bytes.NewBuffer(make([]byte, 0))

	return &VP8Encoder{
		buffer: buffer,
	}
}

func init() {
	registeredEncoders[VP8Codec] newVP8Encoder
}
