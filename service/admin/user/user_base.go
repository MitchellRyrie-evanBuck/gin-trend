package user

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/afl-lxw/gin-trend/global"
	"github.com/afl-lxw/gin-trend/model/admin/system"
	"github.com/gin-gonic/gin"

	"io"
)

type BaseUserConfigService struct {
}

var secretKey = global.TREND_CONFIG.JWT.SigningKey

//type AdminSysUser = system.AdminSysUser

var configUserService = &BaseUserConfigService{}

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
}

type AdminSysUser = system.AdminSysUser

func (b *BaseUserConfigService) UserCreated(c *gin.Context) (err error) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)

		return err
	}
	fmt.Println(input.Password)

	// Check for duplicate username or email
	var existingUser AdminSysUser
	if err := global.TREND_DB.Where("username = ? OR email = ?", input.Username, input.Email).First(&existingUser).Error; err == nil {
		return fmt.Errorf("username or email already exists")
	}

	// Encrypt the password
	encryptedPassword, err := encryptPassword(input.Password, secretKey)
	fmt.Println(encryptedPassword)
	if err != nil {
		return err
	}

	// Save the new user to the database
	user := User{
		Username: input.Username,
		Email:    input.Email,
		Password: encryptedPassword,
	}

	if err := global.TREND_DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
func encryptPassword(password, key string) (string, error) {
	block, err := aes.NewCipher([]byte(createHash(key)))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	return hex.EncodeToString(gcm.Seal(nonce, nonce, []byte(password), nil)), nil
}

func createHash(key string) string {
	hash := sha256.Sum256([]byte(key))
	return hex.EncodeToString(hash[:])
}
