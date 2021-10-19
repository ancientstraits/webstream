package encoders

import "image"

type encoderFactory = func(size image.Point, frameRate int) (Encoder, error)

var registeredEncoders = make(map[VideoCodec]encoderFactory, 2)

type EncoderService struct {
}

func NewEncoderService() Service {
	return &EncoderService{}
}

func (*EncoderService) Supports(codec VideoCodec) bool {
	_, found := registeredEncoders[codec]
	return found
}
