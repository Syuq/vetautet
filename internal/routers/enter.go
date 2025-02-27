package routers

import (
	"github.com/Syuq/go-vetautet-backend-api/internal/routers/manage"
	"github.com/Syuq/go-vetautet-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
