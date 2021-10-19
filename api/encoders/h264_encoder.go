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

func newH264Encoder(size image.Point, frameRate int) (Encoder, error) {
	buffer := bytes.NewBuffer(make([]byte, 0))
	realSize, err := findBestSizeForH264Profile(h264SupportedProfile, size)
	if err != nil {
		return nil, err
	}
	opts := x264.Options{
		Width:     realSize.X,
		Height:    realSize.Y,
		FrameRate: frameRate,
		Tune:      "zerolatency",
		Preset:    "veryfast",
		Profile:   "baseline",
		LogLevel:  x264.LogWarning,
	}
	encoder, err := x264.NewEncoder(buffer, &opts)
	if err != nil {
		return nil, err
	}
	return &H264Encoder{
		buffer:   buffer,
		encoder:  encoder,
		realSize: realSize,
	}, nil
}

const h264SupportedProfile = "3.1"

func (e *H264Encoder) VideoSize() (image.Point, error) {
	return e.realSize, nil
}
