package repo

import (
	"tmi-gin/entity"

	"gorm.io/gorm"
)

type MahasiswaRepository interface {
	UpdateMahasiswa(mahasiswa entity.Mahasiswa) (entity.Mahasiswa, error)
}

type mahasiswaRepo struct {
	connection *gorm.DB
}

func NewMahasiswaRepo(connection *gorm.DB) MahasiswaRepository {
	return &mahasiswaRepo{
		connection: connection,
	}
}

func (c *mahasiswaRepo) UpdateMahasiswa(mahasiswa entity.Mahasiswa) (entity.Mahasiswa, error) {
	c.connection.Save(&mahasiswa)
	c.connection.Preload("User").Find(&mahasiswa)
	return mahasiswa, nil
}
