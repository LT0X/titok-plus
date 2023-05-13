package logic

import (
	"context"
	"strconv"
	"tiktok-plus/common/error/rpcErr"
	"tiktok-plus/common/mq"
	"tiktok-plus/common/utils"

	"tiktok-plus/service/rpc/video/internal/svc"
	"tiktok-plus/service/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteActionLogic) FavoriteAction(in *video.FavoriteActionRequest) (*video.Empty, error) {

	if in.ActionType == 1 {
		//点赞操作
		videoKey := utils.GetFavoriteKey(in.VideoId)
		_, err := l.svcCtx.Redis.
			HSet(l.ctx, videoKey, in.UserId, 1).Result()
		if err != nil {
			return nil, rpcErr.CacheBaseError.WithDetails(err.Error())
		}

		//创建任务队列，视频点赞数缓存加一，异步执行,
		task, err := mq.NewUpdateCacheInfoTask(utils.GetFavoriteCountKey(), strconv.FormatInt(in.UserId, 10), 1)
		if err != nil {
			return nil, rpcErr.MessageQueueError.WithDetails("创建任务队列失败")
		}
		_, err = l.svcCtx.AsynqClient.Enqueue(task)
		if err != nil {
			return nil, rpcErr.MessageQueueError.WithDetails(err.Error())
		}

		//创建任务队列，视频喜爱列表
		task, err = mq.NewUpdateCacheFavoriteList(utils.GetUserLikeSetKey(in.UserId),
			utils.GetUserCancelSetKey(in.UserId), in.VideoId, in.ActionType)
		if err != nil {
			return nil, rpcErr.MessageQueueError.WithDetails("创建任务队列失败")
		}
		_, err = l.svcCtx.AsynqClient.Enqueue(task)
		if err != nil {
			return nil, rpcErr.MessageQueueError.WithDetails(err.Error())
		}

	} else if in.ActionType == 2 {
		//点赞取消操作
		videoKey := utils.GetFavoriteKey(in.VideoId)
		_, err := l.svcCtx.Redis.
			HSet(l.ctx, videoKey, in.UserId, 0).Result()
		if err != nil {
			return nil, rpcErr.CacheBaseError.WithDetails(err.Error())
		}

		//创建任务队列，视频点赞数缓存加一，异步执行,
		task, err := mq.NewUpdateCacheInfoTask(videoKey, strconv.FormatInt(in.UserId, 10), -1)
		if err != nil {
			return nil, rpcErr.MessageQueueError.WithDetails("创建任务队列失败")
		}
		_, err = l.svcCtx.AsynqClient.Enqueue(task)
		if err != nil {
			return nil, rpcErr.MessageQueueError.WithDetails(err.Error())
		}

		//创建任务队列，视频喜爱列表
		task, err = mq.NewUpdateCacheFavoriteList(utils.GetUserLikeSetKey(in.UserId),
			utils.GetUserCancelSetKey(in.UserId), in.VideoId, in.ActionType)
		if err != nil {
			return nil, rpcErr.MessageQueueError.WithDetails("创建任务队列失败")
		}
		_, err = l.svcCtx.AsynqClient.Enqueue(task)
		if err != nil {
			return nil, rpcErr.MessageQueueError.WithDetails(err.Error())
		}
	}
	return &video.Empty{}, nil
}
