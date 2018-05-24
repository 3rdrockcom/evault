package router

import (
	"github.com/epointpayment/evault/app/router/middleware/auth"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// appendMiddleware registers middleware
func (r *Router) appendMiddleware() {
	r.e.Use(middleware.Gzip())
	r.e.Use(middleware.Logger())
	r.e.Use(middleware.Recover())
}

// mwBasicAuth handles the basic authentication for a specific route
func (r *Router) mwBasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(auth.BasicValidator)
}
