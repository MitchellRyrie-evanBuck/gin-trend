package weChat

import (
	"github.com/afl-lxw/gin-trend/global"
	"github.com/afl-lxw/gin-trend/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type WeChart struct {
}

func (t *WeChart) WXCheckSignature(c *gin.Context) {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")
	ok := utils.CheckSignature(signature, timestamp, nonce, global.TREND_CONFIG.WECHAT.Token)
	if !ok {
		log.Println("微信公众号接入校验失败!")
		return
	}

	global.TREND_LOG.Info("微信公众号接入校验成功!")
	global.TREND_LOG.Info(echostr)
	_, _ = c.Writer.WriteString(echostr)

}
