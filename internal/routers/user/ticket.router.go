package user

import (
	"github.com/Syuq/go-vetautet-backend-api/internal/controller/ticket"
	"github.com/gin-gonic/gin"
)

type TicketRouter struct{}

func (tr *TicketRouter) InitTicketRouter(Router *gin.RouterGroup) {
	// public router
	ticketRouterPublic := Router.Group("/ticket")
	{
		// ticketRouterPublic.GET("/search")
		ticketRouterPublic.GET("/item/:id", ticket.TicketItem.GetTicketItemById)
	}
	// private router

}
