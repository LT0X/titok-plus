package video

import (
	"context"

	"tiktok-plus/common/error/apiErr"
	"tiktok-plus/common/utils"

	"tiktok-plus/service/rpc/video/video"

	"tiktok-plus/service/http/internal/svc"
	"tiktok-plus/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedRequest) (resp *types.FeedResponse, err error) {

	logx.WithContext(l.ctx).Infof("视频流: %v", req)
	uid, err := utils.GetUserIDFormToken(req.Token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		uid = 0
	}

	res, err := l.svcCtx.VideoRpc.GetVideoFeed(l.ctx, &video.GetVideoFeedRequest{
		LatestTime: req.LatestTime,
	})

	resp = &types.FeedResponse{
		BaseResponse: types.BaseResponse(apiErr.Success),
		NextTime:     res.VideoList[0].UploadTime,
	}

	videoList := make([]types.Video, 30)
	for _, v := range res.VideoList {
		videoList = append(videoList, types.Video{
			Id:            v.Id,
			UserInfoId:    v.UserInfoId,
			CommentCount:  v.CommentCount,
			FavoriteCount: v.FavoriteCount,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			Title:         v.Title,
		})
	}

	_ = uid

	return resp, nil
}

//
//func (l *FeedLogic) feedInfoByBatch(videoList []*video.VideoInfo, uid int64) (*[]types.Video, error) {
//
//	// 用于 reduce 时保持原来的顺序
//	// likeVideoList 中 videoId 是唯一的, key 选择 videoId, value 是该 video 再 likeVideoList 中原始的位置
//	orderMp := make(map[int]int, len(videoList))
//
//	// mapreduce 并发处理列表请求
//	videoList, err := mr.MapReduce(func(source chan<- interface{}) {
//		for i, v := range videoList {
//			source <- v
//			orderMp[int(v.Id)] = i
//		}
//
//	}, func(item interface{}, writer mr.Writer[types.Video], cancel func(error)) {
//		videoInfo := item.(*video.VideoInfo)
//		// 获取作者信息
//		authorInfo, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &userclient.GetUserByIdRequest{
//			UserId: videoInfo.UserInfoId,
//		})
//
//		if err != nil {
//			logx.WithContext(l.ctx).Errorf("获取作者信息失败: %v", err)
//			cancel(apiErr.ServerInternal)
//			return
//		}
//
//		if uid != 0 {
//			// 获取用户是否关注该作者和点赞视频，
//			isFollowRes, err := l.svcCtx.UserRpc.IsFollow(l.ctx, &userclient.IsFollowRequest{
//				UserId:       UserId,
//				FollowUserId: authorInfo.Id,
//			})
//
//			if err != nil {
//				logx.WithContext(l.ctx).Errorf("获取用户是否关注该作者失败: %v", err)
//				cancel(apiErr.InternalError(l.ctx, err.Error()))
//				return
//			}
//
//
//		} else {
//
//		}
//
//		author := types.User{
//			Id:            authorInfo.Id,
//			Name:          authorInfo.Name,
//			FollowCount:   authorInfo.FollowCount,
//			FollowerCount: authorInfo.FanCount,
//			IsFollow:      isFollowRes.IsFollow,
//		}
//
//		writer.Write(types.Video{
//			Id:            videoInfo.Id,
//			Title:         videoInfo.Title,
//			Author:        author,
//			PlayUrl:       videoInfo.PlayUrl,
//			CoverUrl:      videoInfo.CoverUrl,
//			FavoriteCount: videoInfo.FavoriteCount,
//			CommentCount:  videoInfo.CommentCount,
//			// 这里查询的是用户喜欢的视频列表,无需获取用户是否喜欢
//			IsFavorite: true,
//		})
//
//	}, func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
//		list := make([]types.Video, len(likeVideoList.VideoList))
//		for item := range pipe {
//			videoInfo := item.(types.Video)
//			i, ok := orderMp[int(videoInfo.Id)]
//			if !ok {
//				cancel(apiErr.InternalError(l.ctx, "获取视频在列表中的原始位置失败"))
//				return
//			}
//
//			list[i] = videoInfo
//		}
//
//		writer.Write(list)
//	})
//
//	return nil, err
//
//}
