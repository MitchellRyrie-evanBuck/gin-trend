package response

import "github.com/afl-lxw/gin-trend/model/admin/system"

type SysUserResponse struct {
	User system.AdminSysUser `json:"user"`
}

type LoginResponse struct {
	User      system.AdminSysUser `json:"user"`
	Token     string              `json:"token"`
	ExpiresAt int64               `json:"expiresAt"`
}
