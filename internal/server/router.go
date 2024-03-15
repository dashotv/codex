package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.infratographer.com/x/echox/echozap"
)

func setupRouter(s *Server) {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(echozap.Middleware(s.Logger.Named("router").Desugar()))
	// e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
	// 	Index:      "index.html", // This is the default html page for your SPA
	// 	Browse:     false,
	// 	HTML5:      true,
	// 	Filesystem: http.FS(static.FS),
	// })) // https://echo.labstack.com/docs/middleware/static
	e.HTTPErrorHandler = customHTTPErrorHandler
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(200, "Hello, World!")
	// })

	s.Router = e
}

func customHTTPErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if !ok {
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	code := he.Code
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, map[string]string{"error": "true", "message": he.Error()})
		}
		if err != nil {
			c.Logger().Error(fmt.Errorf("error handling error: %w", err))
		}
	}
}
