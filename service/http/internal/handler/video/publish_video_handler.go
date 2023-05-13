package video

import (
	"bytes"
	"github.com/h2non/filetype"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"tiktok-plus/common/error/apiErr"
	"tiktok-plus/common/utils"
	"tiktok-plus/service/http/internal/logic/video"
	"tiktok-plus/service/http/internal/svc"
	"tiktok-plus/service/http/internal/types"
	"tiktok-plus/service/rpc/user/user"
	video2 "tiktok-plus/service/rpc/video/video"
)

func PublishVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishVideoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取用户id
		userId, err := utils.GetUserIDFormToken(req.Token, svcCtx.Config.Auth.AccessSecret)
		if err != nil {
			httpx.Error(w, apiErr.InvalidToken)
			return
		}

		// 获取文件
		file, fileHeader, err := r.FormFile("data")
		if err != nil {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails(err.Error()))
			return
		}

		defer func(file multipart.File) {
			_ = file.Close()
		}(file)

		// 判断是否为视频
		tmpFile, err := fileHeader.Open()
		if err != nil {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails(err.Error()))
			return
		}
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, tmpFile); err != nil {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails(err.Error()))
			return
		}
		if !filetype.IsVideo(buf.Bytes()) {
			httpx.Error(w, apiErr.FileIsNotVideo)
			return
		}
		if err = tmpFile.Close(); err != nil {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails(err.Error()))
			return
		}

		//获取文件ID
		// 使用uuid重新生成文件名
		fileName := utils.GetUUID() + filepath.Ext(fileHeader.Filename)

		//开始存储
		videoPath := svcCtx.Config.Path.StaticPath + "video/" + fileName
		err = utils.SaveFile(file, videoPath)
		if err = tmpFile.Close(); err != nil {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails(err.Error()))
			return
		}

		//开始获取视频封面并保存
		picturePath := svcCtx.Config.Path.StaticPath + "image/" + utils.GetUUID()
		_, err = utils.GetVideoPicture(videoPath, picturePath, 1)
		if err != nil {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails(err.Error()))
			return
		}

		//开始信息数据库持久化
		_, err = svcCtx.VideoRpc.PublishVideo(r.Context(), &video2.CreateVideoRequest{
			UserInfoId: userId,
			Title:      req.Title,
			CoverUrl:   picturePath + "png",
			PlayUrl:    videoPath,
		})

		if err != nil {
			logx.WithContext(r.Context()).Errorf("PublishVideo rpc error: %v", err)
			httpx.Error(w, apiErr.ServerInternal)
			return
		}

		//同步增长用户视频数

		_, err = svcCtx.UserRpc.AddUserWorkCount(r.Context(), &user.AddUserWorkCountRequest{
			UserInfoId: userId,
		})

		if err != nil {
			logx.WithContext(r.Context()).Errorf("AddUserWorkCount rpc error: %v", err)
			httpx.Error(w, apiErr.ServerInternal)
			return
		}

		l := video.NewPublishVideoLogic(r.Context(), svcCtx)
		resp, err := l.PublishVideo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}

	}
}

func FileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		path := svcCtx.Config.Path.StaticPath
		http.StripPrefix("/static/", http.FileServer(http.Dir(path))).ServeHTTP(w, req)
	}
}
