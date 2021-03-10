package file

import (
	"github.com/lndaquino/segmed-backend/pkg/entity"
)

// FileUsecase struct models a file usecase
type FileUsecase struct {
	repo FileRepo
}

// FileRepo interface defines the methods a file repository must have
type FileRepo interface {
	FindAll() ([]*entity.File, error)
	Update(*entity.File) error
}

// NewFileUsecase returns a FileUsecase instance
func NewFileUsecase(repo FileRepo) *FileUsecase {
	return &FileUsecase{
		repo: repo,
	}
}

// Findall file usecase method handles business rules when retrieving all files from database
func (usecase *FileUsecase) FindAll() ([]*entity.File, error) {
	return usecase.repo.FindAll()
}

// Update file usecase method handles business rules when updating a file record
func (usecase *FileUsecase) Update(id string, status bool) (*entity.File, error) {
	file := &entity.File{
		ID:     id,
		Status: &status,
	}
	if err := usecase.repo.Update(file); err != nil {
		return nil, err
	}
	return file, nil
}
