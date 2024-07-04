package validators

import (
	"errors"
	"fmt"

	"github.com/afl-lxw/gin-trend/dto"
	"github.com/afl-lxw/gin-trend/global"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// validate 实例
var validate = validator.New()

// handleValidationErrors 处理验证错误
func handleValidationErrors(err error) []string {
	if err == nil {
		return nil
	}

	// 使用 errors.As() 进行类型断言，更清晰地处理错误情况
	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		var validationErrors []string
		for _, err := range validationErrs {
			validationErrors = append(validationErrors, generateErrorMessage(err))
		}
		return validationErrors
	}

	// 处理 InvalidValidationError 类型的错误
	var invalidErr *validator.InvalidValidationError
	if errors.As(err, &invalidErr) {
		global.TREND_LOG.Info("user create error", zap.Error(invalidErr))
		return []string{invalidErr.Error()}
	}

	// 如果不是预期的验证错误类型，记录错误并返回一个通用的错误消息
	global.TREND_LOG.Info("user create error", zap.Error(err))
	return []string{"An unexpected error occurred during validation"}
}

// generateErrorMessage 生成更友好的错误消息
func generateErrorMessage(err validator.FieldError) string {
	field := err.Field()
	tag := err.Tag()
	switch tag {
	case "required":
		return fmt.Sprintf("Field '%s' is required", field)
	case "min":
		return fmt.Sprintf("Field '%s' must be greater than or equal to %s", field, err.Param())
	// 可以根据需要添加其他标签的处理
	default:
		return fmt.Sprintf("Field '%s' validation failed on '%s' condition", field, tag)
	}
}

// ValidateCreateUser 验证用户创建请求
func ValidateCreateUser(req *dto.CreateUserRequest) []string {
	return handleValidationErrors(validate.Struct(req))
}
