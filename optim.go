package optim

import (
	"github.com/jayden1228/optim/internal/pkg/audio"
	"github.com/jayden1228/optim/internal/pkg/env"
	"github.com/jayden1228/optim/internal/pkg/image"
	"github.com/jayden1228/optim/internal/pkg/logger"
	"github.com/jayden1228/optim/internal/pkg/video"

	"github.com/gogf/gf/os/gfile"
	"github.com/xfrr/goffmpeg/transcoder"
)

const (
	mp4Suffix        = ".mp4"
	mp3Suffix        = ".mp3"
	pngSuffix        = ".png"
	jpgSuffix        = ".jpg"
	jpegSuffix       = ".jpeg"
)

const (
	audioType = "audio"
	videoType = "video"
	imageType = "image"
	allType = "all"
)

type TranscoderInstance struct {
	Trans *transcoder.Transcoder
}

// NewTranscoderInstance  创建实体
func NewTranscoderInstance() TranscoderInstance{
	if err := env.CheckToolRequired(); err != nil {
		logger.LogE(err)
		panic(err)
	}
	trans :=  new(transcoder.Transcoder)
	return TranscoderInstance{Trans:trans}
}

//  Optim 根据文件类型处理
func (t *TranscoderInstance) Optim(fPath, outPath string) (err error) {
		switch gfile.Ext(fPath) {
		case mp4Suffix:
			err = video.OptimVideoH264(t.Trans, fPath, outPath)
		case mp3Suffix:
		err = audio.OptimAudio(t.Trans, fPath, outPath)
	case pngSuffix, jpegSuffix, jpgSuffix:
		err = image.OptimImage(fPath, outPath)
	}
	return
}
