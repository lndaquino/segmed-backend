package repos

import (
	"github.com/jinzhu/gorm"
	"github.com/lndaquino/segmed-backend/pkg/entity"
	"github.com/lndaquino/segmed-backend/pkg/errors"
)

// FileRepo struct models a file repository
type FileRepo struct {
	db *gorm.DB
}

// NewFileRepo returns a FileRepo instance
func NewFileRepo(db *gorm.DB) *FileRepo {
	return &FileRepo{
		db: db,
	}
}

// FindAll file repo method returns all files entities in database
func (repo *FileRepo) FindAll() ([]*entity.File, error) {
	var files []*entity.File
	if err := repo.db.Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

// Update file repo method updates a file record in database
func (repo *FileRepo) Update(file *entity.File) error {
	req := repo.db.Model(&entity.File{}).Update(&file)
	if req.Error != nil {
		return req.Error
	}
	if req.RowsAffected == 0 {
		return errors.ErrFileNotFound
	}

	return repo.db.Find(&file).Error
}
