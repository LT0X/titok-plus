package handler

import (
	"bytes"
	"douyin/package/util"
	"douyin/response"
	"douyin/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"io"
	"path"
)

func UserRegister(c *fiber.Ctx) error {
	var userService service.UserService
	err := c.QueryParser(&userService)
	if err != nil {
		zap.L().Error(err.Error())
		res := response.UserRegisterOrLogin{
			StatusCode: response.Failed,
			StatusMsg:  response.BadParaRequest,
		}
		c.Status(fiber.StatusOK)
		return c.JSON(res)
	}
	res, err := userService.RegisterService()
	if err != nil {
		res := response.UserRegisterOrLogin{
			StatusCode: response.Failed,
			StatusMsg:  err.Error(),
		}
		c.Status(fiber.StatusOK)
		return c.JSON(res)
	}
	c.Status(fiber.StatusOK)
	return c.JSON(res)
}

func UserLogin(c *fiber.Ctx) error {
	var userService service.UserService
	err := c.QueryParser(&userService)
	if err != nil {
		zap.L().Error(err.Error())
		res := response.UserRegisterOrLogin{
			StatusCode: response.Failed,
			StatusMsg:  response.BadParaRequest,
		}
		c.Status(fiber.StatusOK)
		return c.JSON(res)
	}
	res, err := userService.LoginService()
	if err != nil {
		res := response.UserRegisterOrLogin{
			StatusCode: response.Failed,
			StatusMsg:  err.Error(),
		}
		c.Status(fiber.StatusOK)
		return c.JSON(res)
	}
	c.Status(fiber.StatusOK)
	return c.JSON(res)
}

func UserInfo(c *fiber.Ctx) error {
	var userService service.UserService
	err := c.QueryParser(&userService)
	token := c.Get("token")
	if userService.Token == "" {
		userService.Token = token
	}
	if err != nil {
		zap.L().Error(err.Error())
		res := response.UserRegisterOrLogin{
			StatusCode: response.Failed,
			StatusMsg:  response.BadParaRequest,
		}
		c.Status(fiber.StatusOK)
		return c.JSON(res)
	}
	var userID uint64
	if userService.Token == "" {
		userID = 0
	} else {
		claims, err := util.ParseToken(userService.Token)
		if err != nil {
			res := response.UserRegisterOrLogin{
				StatusCode: response.Failed,
				StatusMsg:  response.WrongToken,
			}
			c.Status(fiber.StatusOK)
			return c.JSON(res)
		}
		userID = claims.UserID
	}
	res, err := userService.InfoService(userID)
	if err != nil {
		res := response.UserRegisterOrLogin{
			StatusCode: response.Failed,
			StatusMsg:  err.Error(),
		}
		c.Status(fiber.StatusOK)
		return c.JSON(res)
	}
	c.Status(fiber.StatusOK)
	return c.JSON(res)
}

func UpdateUserInfo(c *fiber.Ctx) error {
	var userService service.UserService
	err := c.QueryParser(&userService)
	token := c.Get("token")
	if userService.Token == "" {
		userService.Token = token
	}
	if userService.Token == "" {
		token = c.FormValue("token")
		userService.Token = token
	}

	if err != nil {
		zap.L().Error(err.Error())
		res := response.UserRegisterOrLogin{
			StatusCode: response.Failed,
			StatusMsg:  response.BadParaRequest,
		}
		c.Status(fiber.StatusOK)
		return c.JSON(res)
	}
	var userID uint64
	if userService.Token == "" {
		userID = 0
	} else {
		claims, err := util.ParseToken(userService.Token)
		if err != nil {
			res := response.UserRegisterOrLogin{
				StatusCode: response.Failed,
				StatusMsg:  response.WrongToken,
			}
			c.Status(fiber.StatusOK)
			return c.JSON(res)
		}
		userID = claims.UserID
	}
	username := c.FormValue("username")
	signature := c.FormValue("signature")
	fileHeader, err := c.FormFile("avatar")
	if err != nil {
		zap.L().Error(err.Error())
		res := response.CommonResponse{
			StatusCode: response.Failed,
			StatusMsg:  response.FileFormatError,
		}
		return c.JSON(res)
	}
	extension := path.Ext(fileHeader.Filename)
	// 检查文件后缀是不是png,jpg 大小在上传的时候会限制30MB
	if extension != ".png" && extension != ".jpg" && extension != ".jpeg" {
		zap.L().Error(err.Error())
		res := response.CommonResponse{
			StatusCode: response.Failed,
			StatusMsg:  response.FileFormatError,
		}
		return c.JSON(res)
	}

	//if !strings.HasSuffix(fileHeader.Filename, ".png") {
	//	zap.L().Error(err.Error())
	//	res := response.CommonResponse{
	//		StatusCode: response.Failed,
	//		StatusMsg:  response.FileFormatError,
	//	}
	//	return c.JSON(res)
	//}

	zap.L().Info("PublishAction Filename:" + fileHeader.Filename)

	file, err := fileHeader.Open()
	if err != nil {
		zap.L().Error(err.Error())
		res := response.CommonResponse{
			StatusCode: response.Failed,
			StatusMsg:  response.FileFormatError,
		}
		return c.JSON(res)
	}
	defer file.Close()
	// 将文件转化为字节流
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		zap.L().Error(err.Error())
		res := response.CommonResponse{
			StatusCode: response.Failed,
			StatusMsg:  response.FileFormatError,
		}
		return c.JSON(res)
	}

	res, err := userService.UpdateInfoService(userID, username, signature, extension, buf)
	if err != nil {
		res := response.UserRegisterOrLogin{
			StatusCode: response.Failed,
			StatusMsg:  err.Error(),
		}
		c.Status(fiber.StatusOK)
		return c.JSON(res)
	}
	c.Status(fiber.StatusOK)
	return c.JSON(res)
}

func SearchUser(c *fiber.Ctx) error {
	var userService service.UserService
	err := c.QueryParser(&userService)
	token := c.Get("token")
	if userService.Token == "" {
		userService.Token = token
	}
	if err != nil {
		zap.L().Error(err.Error())
		res := response.UserRegisterOrLogin{
			StatusCode: response.Failed,
			StatusMsg:  response.BadParaRequest,
		}
		c.Status(fiber.StatusOK)
		return c.JSON(res)
	}
	var userID uint64
	//if userService.Token == "" {
	//	userID = 0
	//} else {
	//	claims, err := util.ParseToken(userService.Token)
	//	if err != nil {
	//		res := response.UserRegisterOrLogin{
	//			StatusCode: response.Failed,
	//			StatusMsg:  response.WrongToken,
	//		}
	//		c.Status(fiber.StatusOK)
	//		return c.JSON(res)
	//	}
	//	userID = claims.UserID
	//}

	username := c.Query("username")

	res, err := userService.SearchUserService(int64(userID), username)
	if err != nil {
		res := response.UserRegisterOrLogin{
			StatusCode: response.Failed,
			StatusMsg:  err.Error(),
		}
		c.Status(fiber.StatusOK)
		return c.JSON(res)
	}
	c.Status(fiber.StatusOK)
	return c.JSON(res)
}
