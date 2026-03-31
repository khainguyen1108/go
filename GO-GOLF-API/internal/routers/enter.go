package routers

import (
	"GO-GOLF-API/internal/routers/manage"
	"GO-GOLF-API/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
