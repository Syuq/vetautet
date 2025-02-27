package initialize

import (
	"github.com/Syuq/go-vetautet-backend-api/global"
	"github.com/Syuq/go-vetautet-backend-api/internal/database"
	"github.com/Syuq/go-vetautet-backend-api/internal/service"
	"github.com/Syuq/go-vetautet-backend-api/internal/service/impl"
)

func InitServiceInterface() {
	queries := database.New(global.Mdbc)
	// User Service Interface
	service.InitUserLogin(impl.NewUserLoginImpl(queries))
	// Ticker Service Interface
	service.InitTicketItem(impl.NewTicketItemImpl(queries))
}
