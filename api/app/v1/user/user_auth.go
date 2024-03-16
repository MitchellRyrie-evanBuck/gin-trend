package user

import "github.com/gin-gonic/gin"

type BaseUserOAuthToken struct {
}

var tokenData struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
}

// OAuthToken @Title OAuth获取用户信息
func (t *BaseUserOAuthToken) OAuthToken(c *gin.Context) {

}
