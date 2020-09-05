package video

import (
	"log"

	"github.com/xfrr/goffmpeg/transcoder"
)

func OptimVideoH264(trans *transcoder.Transcoder, inputPath, outputPath string) error {
	err := trans.Initialize(inputPath, outputPath)

	if err != nil {
		return err
	}

	trans.MediaFile().SetVideoCodec("libx264")

	log.Println("start compress: ", inputPath)

	done := trans.Run(true)

	log.Println("finish compress: ", inputPath)

	//progress := trans.Output()
	//
	//for msg := range progress {
	//	log.Println(msg)
	//}

	err = <-done

	return err
}
