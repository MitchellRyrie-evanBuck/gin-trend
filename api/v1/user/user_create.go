package user

import "github.com/gin-gonic/gin"

func (b *BaseUserApi) UserCreatedAPIFn(c *gin.Context) {
	err := configUserService.UserCreated(c)
	if err != nil {

	}
}
