package request

type Login struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Iphone   string `json:"Iphone"`
	Code     string `json:"code"`
}
