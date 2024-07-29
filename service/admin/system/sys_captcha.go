package system

import (
	"bytes"
	"context"
	"encoding/base64"
	"github.com/afl-lxw/gin-trend/global"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

type BaseSystemCaptchaService struct {
}

var (
	rdb *redis.Client
	ctx = context.Background()
)

// CaptchaResponse 验证码响应结构
type CaptchaResponse struct {
	Base64 string `json:"base64"`
	ID     string `json:"id"`
}

// GenerateCaptchaHandler 生成验证码接口
func (t *BaseSystemCaptchaService) GenerateCaptchaHandler(c *gin.Context) (data CaptchaResponse, err error) {
	// 生成一个新的验证码ID
	id := captcha.New()

	// 获取验证码图像
	var buf bytes.Buffer
	captcha.WriteImage(&buf, id, global.TREND_CONFIG.Captcha.ImgWidth,
		global.TREND_CONFIG.Captcha.ImgHeight,
	)

	// 将验证码图像转换为Base64编码
	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())

	err = global.TREND_REDIS.Set(ctx, id, base64Str, time.Minute).Err()
	if err != nil {
		return CaptchaResponse{}, err
	}

	// 创建响应
	response := CaptchaResponse{
		Base64: "data:image/png;base64," + base64Str,
		ID:     id,
	}
	return response, nil
}

// verifyCaptchaHandler 验证验证码接口
// 验证验证码的处理函数
func verifyCaptchaHandler(c *gin.Context) {
	id := c.Query("id")
	value := c.Query("value")

	if id == "" || value == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id and value are required"})
		return
	}

	// 从Redis中获取验证码
	_, err := global.TREND_REDIS.Get(ctx, id).Result()
	if err == redis.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired captcha"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve captcha"})
		return
	}

	// 验证验证码
	if captcha.VerifyString(id, value) {
		c.JSON(http.StatusOK, gin.H{"message": "Captcha verified successfully!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid captcha"})
	}
}
