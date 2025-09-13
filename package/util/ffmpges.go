package util

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"go.uber.org/zap"
	"os"
)

func GetSnapshot(videoPath, imageName string, frameNum int) (ImagePath string, err error) {
	snapshotPath := "C:\\work\\GoWord\\v1-clip\\DikTok\\douyinVideo\\coverurl\\" + imageName
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		zap.L().Error("生成缩略图失败：" + err.Error())
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		zap.L().Error("生成缩略图失败：" + err.Error())
		return "", err
	}

	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		zap.L().Error("生成缩略图失败：" + err.Error())
		return "", err
	}

	//imgPath := snapshotPath + ".png"
	coverurl := "http://127.0.0.1:8000/static/coverurl/" + imageName + ".png"

	return coverurl, nil
}
