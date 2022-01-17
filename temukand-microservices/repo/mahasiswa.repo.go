package repo

import (
	"tmi-gin/entity"

	"gorm.io/gorm"
)

type MahasiswaRepository interface {
	All(userID string) ([]entity.Mahasiswa, error)
	InsertMahasiswa(mahasiswa entity.Mahasiswa) (entity.Mahasiswa, error)
	UpdateMahasiswa(mahasiswa entity.Mahasiswa) (entity.Mahasiswa, error)
	DeleteMahasiswa(mahasiswaID string) error
	FindOneMahasiswaByID(ID string) (entity.Mahasiswa, error)
	FindAllMahasiswa(userID string) ([]entity.Mahasiswa, error)
}

type mahasiswaRepo struct {
	connection *gorm.DB
}

func NewMahasiswaRepo(connection *gorm.DB) MahasiswaRepository {
	return &mahasiswaRepo{
		connection: connection,
	}
}

func (c *mahasiswaRepo) All(userID string) ([]entity.Mahasiswa, error) {
	mahasiswas := []entity.Mahasiswa{}
	c.connection.Preload("User").Where("user_id = ?", userID).Find(&mahasiswas)
	return mahasiswas, nil
}

func (c *mahasiswaRepo) InsertMahasiswa(mahasiswa entity.Mahasiswa) (entity.Mahasiswa, error) {
	c.connection.Save(&mahasiswa)
	c.connection.Preload("User").Find(&mahasiswa)
	return mahasiswa, nil
}

func (c *mahasiswaRepo) UpdateMahasiswa(mahasiswa entity.Mahasiswa) (entity.Mahasiswa, error) {
	c.connection.Save(&mahasiswa)
	c.connection.Preload("User").Find(&mahasiswa)
	return mahasiswa, nil
}

func (c *mahasiswaRepo) FindOneMahasiswaByID(mahasiswaID string) (entity.Mahasiswa, error) {
	var mahasiswa entity.Mahasiswa
	res := c.connection.Preload("User").Where("id = ?", mahasiswaID).Take(&mahasiswa)
	if res.Error != nil {
		return mahasiswa, res.Error
	}
	return mahasiswa, nil
}

func (c *mahasiswaRepo) FindAllMahasiswa(userID string) ([]entity.Mahasiswa, error) {
	mahasiswas := []entity.Mahasiswa{}
	c.connection.Where("user_id = ?", userID).Find(&mahasiswas)
	return mahasiswas, nil
}

func (c *mahasiswaRepo) DeleteMahasiswa(mahasiswaID string) error {
	var mahasiswa entity.Mahasiswa
	res := c.connection.Preload("User").Where("id = ?", mahasiswaID).Take(&mahasiswa)
	if res.Error != nil {
		return res.Error
	}
	c.connection.Delete(&mahasiswa)
	return nil
}
