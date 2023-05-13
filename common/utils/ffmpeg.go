package utils

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"
)

// GetVideoPicture 调用ffmpeg为视频截取封面
func GetVideoPicture(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {

	buf := bytes.NewBuffer(nil)
	err = ffmpeg_go.Input(videoPath).
		Filter("select", ffmpeg_go.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg_go.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).Run()
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}
	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}
	names := strings.Split(snapshotPath, "\\")
	snapshotName = names[len(names)-1] + ".png"
	return
}

func SaveFile(file multipart.File, path string) error {
	defer file.Close()

	// 创建一个新文件
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// 将multipart.File中的数据写入文件中
	_, err = io.Copy(f, file)
	return err
}
