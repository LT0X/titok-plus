package llm

import (
	"context"
	"douyin/database"
	"douyin/model"
	"douyin/rpc/contact/contact"
	"strconv"

	"go.uber.org/zap"
)

var (
	// chatgpt
	ChatGPTAvatar = "http://127.0.0.1:8000/static/avater/gpt.jpg"
	ChatGPTName   = "ChatGPT"
	ChatGPTID     = uint64(1)
)

func SendToChatGPT(userID uint64, content string) error {
	// 先将消息写入数据库

	//err := database.CreateMessage(userID, ChatGPTID, content)
	res, err := database.RPC.ContactRpc.CreateMessage(context.TODO(), &contact.CreateMessageRequest{
		UserID:   userID,
		ToUserID: ChatGPTID,
		Content:  content,
	})
	if err != nil || res.Reply != 1 {
		return err
	}

	if err != nil {
		return err
	}
	go requestToChatGPT(userID, content)
	return nil
}

func requestToChatGPT(userID uint64, content string) {
	ans := RequestToSparkAPI(content)
	if ans == "" {
		return
	}

	//err := database.CreateMessage(ChatGPTID, userID, ans)
	res, err := database.RPC.ContactRpc.CreateMessage(context.TODO(), &contact.CreateMessageRequest{
		UserID:   ChatGPTID,
		ToUserID: userID,
		Content:  ans,
	})

	if err != nil || res.Reply != 1 {
		zap.L().Error(err.Error() + " reply" + strconv.FormatUint(res.Reply, 64))
	}
}

// 将chatgpt注册为用户
func RegisterChatGPT() {
	user := &model.User{
		ID:       ChatGPTID,
		Username: ChatGPTName,
		Avatar:   ChatGPTAvatar,
	}
	_, err := database.CreateUser(user)
	if err != nil {
		zap.L().Info("ChatGPT已写入user表")
	}
}
