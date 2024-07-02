package auth

import "github.com/gin-gonic/gin"

type BaseGoogleApi struct {
}

func (t *BaseGoogleApi) VerificationGoogleCode(c *gin.Context) {
	data, err := authGoogleService.HandleGoogleLogin("123", c)
	if err != nil {

	}
	println(data)
}
