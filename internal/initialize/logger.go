package initialize

import (
	"github.com/Syuq/go-vetautet-backend-api/global"
	"github.com/Syuq/go-vetautet-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
