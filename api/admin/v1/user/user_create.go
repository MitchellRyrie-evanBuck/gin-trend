package user

import "github.com/gin-gonic/gin"

func (b *BaseUserApi) UserCreatedAPIFn(c *gin.Context) {
	err := configUserService.UserCreated(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User created successfully"})
}
