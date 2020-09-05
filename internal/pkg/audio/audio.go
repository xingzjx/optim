package audio

import (
	"log"

	"github.com/xfrr/goffmpeg/transcoder"
)

const (
	audioBitRate = "160"
)

func OptimAudio(trans *transcoder.Transcoder, inputPath, outputPath string) error {
	err := trans.Initialize(inputPath, outputPath)

	if err != nil {
		return err
	}

	trans.MediaFile().SetAudioBitRate(audioBitRate)

	log.Println("start compress: ", inputPath)

	done := trans.Run(true)

	log.Println("finish compress: ", inputPath)
	//
	//progress := trans.Output()
	//
	//for msg := range progress {
	//	log.Println(msg)
	//}

	err = <-done

	return err
}
