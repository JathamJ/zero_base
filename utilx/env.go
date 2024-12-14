package utilx

import (
	"github.com/JathamJ/zero_base/constantx"
	"os"
)

// IsProd 是否生产环境
func IsProd() bool {
	env := os.Getenv("APP_ENV")
	return env == constantx.Production
}

// IsDev 是否开发环境
func IsDev() bool {
	env := os.Getenv("APP_ENV")
	return env == constantx.Development || env == ""
}
