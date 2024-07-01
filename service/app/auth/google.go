package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GoogleAuthService struct {
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	// 获取前端发送的授权码
	code := r.FormValue("code")

	// 发送 HTTP 请求到 Google 验证端点
	resp, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=%s", code))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// 解析响应
	var tokenInfo map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&tokenInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 检查验证结果
	if resp.StatusCode == http.StatusOK {
		// 验证成功，处理用户信息
		// 例如，从 tokenInfo 中获取用户的唯一标识符 tokenInfo["sub"]
		// 将用户信息存储在数据库中或者生成用户会话等
		// 返回成功响应给前端
		json.NewEncoder(w).Encode(tokenInfo)
	} else {
		// 验证失败，返回错误信息给前端
		http.Error(w, "Google authentication failed", resp.StatusCode)
	}
}
