package handler

import (
	"mindarc/mindarc-reporting/view"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type MainHandler struct{}

func (h MainHandler) HandleMain(c echo.Context) error {
	return render(c, view.Base())
}

func (h MainHandler) HandleTest(c echo.Context) error {
	return render(c, view.AnotherOne())
}

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
