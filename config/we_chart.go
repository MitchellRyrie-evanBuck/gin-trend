package config

type WECHAT struct {
	AssetToken string `mapstructure:"asset-token" json:"asset-token" yaml:"asset-token"` // 收件人:多个以英文逗号分隔 例：a@qq.com b@qq.com 正式开发中请把此项目作为参数使用
	Signature  string `mapstructure:"signature" json:"signature" yaml:"signature"`       // 发件人  你自己要发邮件的邮箱
	Token      string `mapstructure:"token" json:"token" yaml:"token"`                   // 服务器地址 例如 smtp.qq.com  请前往QQ或者你要发邮件的邮箱查看其smtp协议
}
