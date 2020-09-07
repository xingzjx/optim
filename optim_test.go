package optim

import (
	"testing"
)

func TestOptimAllType(t *testing.T) {

	instance := NewTranscoderInstance()

	instance.Optim("./example/demo.mp3", "./output/demo.mp3")
	instance.Optim("./example/demo.jpg", "./output/demo.jpg")
	instance.Optim("./example/demo.mp4", "./output/demo.mp4")

}