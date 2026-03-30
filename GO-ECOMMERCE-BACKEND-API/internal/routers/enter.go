package routers

import (
	"GO-ECOMMERCE-BACKEND-API/internal/routers/manage"
	"GO-ECOMMERCE-BACKEND-API/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
