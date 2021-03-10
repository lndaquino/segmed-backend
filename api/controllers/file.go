package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/lndaquino/segmed-backend/pkg/entity"
	"github.com/lndaquino/segmed-backend/pkg/errors"
	"github.com/lndaquino/segmed-backend/utils"
)

// FileController struct models a controller for file routes
type FileController struct {
	usecase FileUsecase
}

// FileUsecase interface defines the file usecase methods
type FileUsecase interface {
	FindAll() ([]*entity.File, error)
	Update(string, bool) (*entity.File, error)
}

// NewFileController returns a file controller instance
func NewFileController(usecase FileUsecase) *FileController {
	return &FileController{
		usecase: usecase,
	}
}

// FindAll file controller method handles file routes call asking for all files
func (ctl *FileController) FindAll(c echo.Context) error {
	files, err := ctl.usecase.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, files)
}

// Update file controller method handles file status update and return the updated file or error
func (ctl FileController) Update(c echo.Context) error {
	fileID := c.Param("file_id")
	statusParam := c.QueryParam("status")
	if _, err := uuid.Parse(fileID); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"error": "file_id path param must be uuid type",
		})
	}

	if !(statusParam == "true" || statusParam == "false") {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "status query param must be true or false",
		})
	}

	status, err := strconv.ParseBool(statusParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	updatedFile, err := ctl.usecase.Update(fileID, status)
	if err != nil {
		switch err {
		case errors.ErrFileNotFound:
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": err.Error(),
			})
		default:
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
		}
	}

	response := new(fileUpdateResponse)
	if err := utils.ParseStruct(updatedFile, response); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response)
}

type fileUpdateResponse struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`
	Status    *bool     `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}
