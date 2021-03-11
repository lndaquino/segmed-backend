package repos

import (
	"time"

	"github.com/lndaquino/segmed-backend/pkg/entity"
	"github.com/lndaquino/segmed-backend/pkg/errors"
)

// FileRepo struct models a file repository
type FileInMemoryRepo struct {
	m map[string]*entity.File
}

// NewFileRepo returns a FileRepo instance
func NewFileInMemoryRepo() *FileInMemoryRepo {
	now := time.Now()

	file1 := &entity.File{"730449a9-aafc-48ac-9fd0-e8b7c5e29df4", "https://segmed.blob.core.windows.net/segmed/cdc-AY9uFreZrcw-unsplash.jpg", isFalse(), now}
	file2 := &entity.File{"9381abd3-8163-4e4b-b638-2587ec907fba", "https://segmed.blob.core.windows.net/segmed/cdc-vt7iAyiwpf0-unsplash.jpg", isFalse(), now}
	file3 := &entity.File{"d36e1dfb-9a40-40e7-8aaa-614430f93485", "https://segmed.blob.core.windows.net/segmed/dennis-klicker-J4HwEwZtIs8-unsplash.jpg", isFalse(), now}
	file4 := &entity.File{"51b6dc37-cb91-4221-95de-300b3ad554a1", "https://segmed.blob.core.windows.net/segmed/hh-e-VeF3uWcH4L4-unsplash.jpg", isFalse(), now}
	file5 := &entity.File{"2e12f69e-1625-4cc4-9495-d4f4e8d4ba2e", "https://segmed.blob.core.windows.net/segmed/owen-beard-DK8jXx1B-1c-unsplash.jpg", isFalse(), now}
	file6 := &entity.File{"70af5f08-7d08-4be8-8c6d-502eebf86b59", "https://segmed.blob.core.windows.net/segmed/sharon-mccutcheon-OjxToQpMfW4-unsplash.jpg", isFalse(), now}
	file7 := &entity.File{"619d8a74-6cd1-45d2-aa37-53b2c9f5d920", "https://segmed.blob.core.windows.net/segmed/thula-na-pE_gFtC15mA-unsplash.jpg", isFalse(), now}
	file8 := &entity.File{"6f3749df-a762-44de-8385-098e2c7aa6b6", "https://segmed.blob.core.windows.net/segmed/anne-nygard-_W94Eb1iNYc-unsplash.jpg", isFalse(), now}
	file9 := &entity.File{"82f485cb-e9d8-41f9-b02a-15e4184603ec", "https://segmed.blob.core.windows.net/segmed/austrian-national-library-sOPnjqLjyY8-unsplash.jpg", isFalse(), now}
	file10 := &entity.File{"17cb7364-408d-41b6-abd2-f477326088e9", "https://segmed.blob.core.windows.net/segmed/robina-weermeijer-Tmkwl7EjVtE-unsplash.jpg", isFalse(), now}

	return &FileInMemoryRepo{map[string]*entity.File{
		file1.ID:  file1,
		file2.ID:  file2,
		file3.ID:  file3,
		file4.ID:  file4,
		file5.ID:  file5,
		file6.ID:  file6,
		file7.ID:  file7,
		file8.ID:  file8,
		file9.ID:  file9,
		file10.ID: file10,
	}}
}

// FindAll file repo method returns all files entities in database
func (repo *FileInMemoryRepo) FindAll() ([]*entity.File, error) {
	var files []*entity.File
	for _, file := range repo.m {
		files = append(files, file)
	}
	return files, nil
}

// Update file repo method updates a file record in database
func (repo *FileInMemoryRepo) Update(file *entity.File) error {
	if repo.m[file.ID] == nil {
		return errors.ErrFileNotFound
	}

	updatedFile := repo.m[file.ID]
	if *file.Status != *updatedFile.Status {
		*updatedFile.Status = *file.Status
	}
	if file.URL != "" {
		updatedFile.URL = file.URL
	}
	updatedFile.UpdatedAt = time.Now()

	*repo.m[file.ID] = *updatedFile
	*file = *updatedFile

	return nil
}

func isFalse() *bool {
	isFalse := false
	return &isFalse
}
