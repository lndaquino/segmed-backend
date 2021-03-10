package routers

import (
	"github.com/labstack/echo"
	"github.com/lndaquino/segmed-backend/api/controllers"
)

// SystemRoutes struct models a system level router
type SystemRoutes struct {
	fileController *controllers.FileController
}

// NewSystemRoutes returns a SystemRoutes instance
func NewSystemRoutes(fileController *controllers.FileController) *SystemRoutes {
	return &SystemRoutes{
		fileController: fileController,
	}
}

// MakeControllers defines the system routes and return a wrapped controller
func (routes *SystemRoutes) MakeControllers() *echo.Echo {
	e := echo.New()

	e.GET("/", routes.fileController.FindAll)
	e.PATCH("/:file_id/", routes.fileController.Update)

	return e
}
