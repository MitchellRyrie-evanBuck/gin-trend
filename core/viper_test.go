package core_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/afl-lxw/gin-trend/core"
	"github.com/afl-lxw/gin-trend/global"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestViper_DefaultConfig(t *testing.T) {
	// 设置 gin 模式
	gin.SetMode(gin.DebugMode)

	// 创建临时配置文件
	tmpDir, err := ioutil.TempDir("", "viper_test")
	if err != nil {
		t.Fatalf("无法创建临时目录: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	configContent := []byte(`
Trend:
  Name: "TestDefault"
  AutoCode:
    Root: "/tmp/default"
`)
	tmpFile := filepath.Join(tmpDir, "config.yaml")
	if err := ioutil.WriteFile(tmpFile, configContent, 0644); err != nil {
		t.Fatalf("无法写入临时文件: %v", err)
	}

	// 设置环境变量为临时配置文件路径
	os.Setenv("GIN_TREND_CONFIG", tmpFile)
	defer os.Unsetenv("GIN_TREND_CONFIG")

	// 调用 Viper 函数初始化配置
	v := core.Viper()
	assert.NotNil(t, v)
	assert.Equal(t, "TestDefault", global.TREND_CONFIG.Setting.Name)
	assert.Equal(t, "/tmp/default", global.TREND_CONFIG.AutoCode.Root)
}

func TestViper_WithPath(t *testing.T) {
	// 创建临时配置文件
	tmpDir, err := ioutil.TempDir("", "viper_test")
	if err != nil {
		t.Fatalf("无法创建临时目录: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	configContent := []byte(`
Trend:
  Name: "TestWithPath"
  AutoCode:
    Root: "/tmp/path"
`)
	tmpFile := filepath.Join(tmpDir, "config_path.yaml")
	if err := ioutil.WriteFile(tmpFile, configContent, 0644); err != nil {
		t.Fatalf("无法写入临时文件: %v", err)
	}

	// 调用 Viper 函数，传递临时配置文件路径
	v := core.Viper(tmpFile)
	assert.NotNil(t, v)
	assert.Equal(t, "TestWithPath", global.TREND_CONFIG.Setting.Name)
	assert.Equal(t, "/tmp/path", global.TREND_CONFIG.AutoCode.Root)
}
