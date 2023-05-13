package logic

import (
	"context"
	"strconv"
	"tiktok-plus/common/error/rpcErr"
	"tiktok-plus/common/utils"

	"tiktok-plus/service/rpc/video/internal/svc"
	"tiktok-plus/service/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	PopularVideoStandard = 1000 // 拥有超过 1000 个赞或 1000 个评论的视频成为热门视频，有特殊处理
)

type GetVideoFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoFeedLogic {
	return &GetVideoFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoFeedLogic) GetVideoFeed(in *video.GetVideoFeedRequest) (*video.GetVideoFeedResponse, error) {

	maxNum := 30
	DB := l.svcCtx.DBList.DB
	videoList := make([]*video.VideoInfo, 30)
	err := DB.Table("titok_plus_video").Order("upload_time Desc").
		Limit(maxNum).Where("upload_time < ?", in.LatestTime).
		Find(&videoList).Error
	if err != nil {
		return nil, rpcErr.DataBaseError.WithDetails(err.Error())
	}

	//为热门视频进行缓存
	for _, v := range videoList {
		if IsPopularVideo(v) {
			//查看是否有缓存
			if l.svcCtx.Redis.HExists(l.ctx, utils.GetPopVideoKey(), strconv.FormatInt(v.Id, 10)).Val() == false {
				//不在缓存，添加进入缓存
				err = l.svcCtx.Redis.
					HSet(l.ctx, utils.GetPopVideoKey(), strconv.FormatInt(v.Id, 10), "1").Err()
				if err != nil {
					return nil, rpcErr.CacheBaseError.WithDetails(err.Error())
				}
			}
		}
	}
	return &video.GetVideoFeedResponse{
		VideoList: videoList,
	}, nil
}

func IsPopularVideo(v *video.VideoInfo) bool {
	return v.FavoriteCount > PopularVideoStandard || v.CommentCount > PopularVideoStandard
}
