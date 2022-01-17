package service

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"tmi-gin/dto"
	"tmi-gin/entity"
	"tmi-gin/repo"

	"github.com/mashingan/smapping"

	_mahasiswa "tmi-gin/service/mahasiswa"
)

type MahasiswaService interface {
	All(userID string) (*[]_mahasiswa.MahasiswaResponse, error)
	CreateMahasiswa(mahasiswaRequest dto.CreateMahasiswaRequest, userID string) (*_mahasiswa.MahasiswaResponse, error)
	UpdateMahasiswa(updateMahasiswaRequest dto.UpdateMahasiswaRequest, userID string) (*_mahasiswa.MahasiswaResponse, error)
	FindOneMahasiswaByID(mahasiswaID string) (*_mahasiswa.MahasiswaResponse, error)
	DeleteMahasiswa(mahasiswaID string, userID string) error
}

type mahasiswaService struct {
	mahasiswaRepo repo.MahasiswaRepository
}

func NewMahasiswaService(mahasiswaRepo repo.MahasiswaRepository) MahasiswaService {
	return &mahasiswaService{
		mahasiswaRepo: mahasiswaRepo,
	}
}

func (c *mahasiswaService) All(userID string) (*[]_mahasiswa.MahasiswaResponse, error) {
	mahasiswas, err := c.mahasiswaRepo.All(userID)
	if err != nil {
		return nil, err
	}

	prods := _mahasiswa.NewMahasiswaArrayResponse(mahasiswas)
	return &prods, nil
}

func (c *mahasiswaService) CreateMahasiswa(mahasiswaRequest dto.CreateMahasiswaRequest, userID string) (*_mahasiswa.MahasiswaResponse, error) {
	mahasiswa := entity.Mahasiswa{}
	err := smapping.FillStruct(&mahasiswa, smapping.MapFields(&mahasiswaRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	id, _ := strconv.ParseInt(userID, 0, 64)
	mahasiswa.UserID = id
	p, err := c.mahasiswaRepo.InsertMahasiswa(mahasiswa)
	if err != nil {
		return nil, err
	}

	res := _mahasiswa.NewMahasiswaResponse(p)
	return &res, nil
}

func (c *mahasiswaService) FindOneMahasiswaByID(mahasiswaID string) (*_mahasiswa.MahasiswaResponse, error) {
	mahasiswa, err := c.mahasiswaRepo.FindOneMahasiswaByID(mahasiswaID)

	if err != nil {
		return nil, err
	}

	res := _mahasiswa.NewMahasiswaResponse(mahasiswa)
	return &res, nil
}

func (c *mahasiswaService) UpdateMahasiswa(updateMahasiswaRequest dto.UpdateMahasiswaRequest, userID string) (*_mahasiswa.MahasiswaResponse, error) {
	mahasiswa, err := c.mahasiswaRepo.FindOneMahasiswaByID(fmt.Sprintf("%d", updateMahasiswaRequest.ID))
	if err != nil {
		return nil, err
	}

	uid, _ := strconv.ParseInt(userID, 0, 64)
	if mahasiswa.UserID != uid {
		return nil, errors.New("Ini bukan milik anda")
	}

	mahasiswa = entity.Mahasiswa{}
	err = smapping.FillStruct(&mahasiswa, smapping.MapFields(&updateMahasiswaRequest))

	if err != nil {
		return nil, err
	}

	mahasiswa.UserID = uid
	mahasiswa, err = c.mahasiswaRepo.UpdateMahasiswa(mahasiswa)

	if err != nil {
		return nil, err
	}

	res := _mahasiswa.NewMahasiswaResponse(mahasiswa)
	return &res, nil
}

func (c *mahasiswaService) DeleteMahasiswa(mahasiswaID string, userID string) error {
	mahasiswa, err := c.mahasiswaRepo.FindOneMahasiswaByID(mahasiswaID)
	if err != nil {
		return err
	}

	if fmt.Sprintf("%d", mahasiswa.UserID) != userID {
		return errors.New("Ini bukan milik anda")
	}

	c.mahasiswaRepo.DeleteMahasiswa(mahasiswaID)
	return nil

}
