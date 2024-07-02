package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GoogleAuthService struct {
}

func (t *GoogleAuthService) HandleGoogleLogin(code string, c *gin.Context) (r *http.Response, err error) {
	// 获取前端发送的授权码

	// 发送 HTTP 请求到 Google 验证端点
	resp, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=%s", code))

	if err != nil {
	}

	return resp, err

}
