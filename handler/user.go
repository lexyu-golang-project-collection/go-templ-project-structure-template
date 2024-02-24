package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/lexyu-golang-project-collection/go-templ-project-structure-template/model"
	"github.com/lexyu-golang-project-collection/go-templ-project-structure-template/view/user"
)

type UserHandler struct {
}

func (uh *UserHandler) UserHandlerShow(c echo.Context) error {

	u := model.User{
		Email: "test@gmail.com",
	}

	return render(c, user.Show(u))

	// return user.Show().Render(c.Request().Context(), c.Response())

	// return nil
}
