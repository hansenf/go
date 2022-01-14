package service

import (
	"errors"
	"strconv"

	"tmi-gin/dto"
	"tmi-gin/entity"
	"tmi-gin/repo"

	"github.com/mashingan/smapping"

	_mahasiswa "tmi-gin/service/mahasiswa"
)

type MahasiswaService interface {
	UpdateMahasiswa(updateMahasiswaRequest dto.UpdateMahasiswaRequest, userID string) (*_mahasiswa.MahasiswaResponse, error)
}

type mahasiswaService struct {
	mahasiswaRepo repo.MahasiswaRepository
}

func NewMahasiswaService(mahasiswaRepo repo.MahasiswaRepository) MahasiswaService {
	return &mahasiswaService{
		mahasiswaRepo: mahasiswaRepo,
	}
}

func (c *mahasiswaService) UpdateMahasiswa(updateMahasiswaRequest dto.UpdateMahasiswaRequest, userID string) (*_mahasiswa.MahasiswaResponse, error) {
	mahasiswa, err := c.mahasiswaRepo.UpdateMahasiswa(entity.Mahasiswa{})
	if err != nil {
		return nil, err
	}

	uid, _ := strconv.ParseInt(userID, 0, 64)
	if mahasiswa.User.ID != uid {
		return nil, errors.New("Ini bukan id anda")
	}

	mahasiswa = entity.Mahasiswa{}
	err = smapping.FillStruct(&mahasiswa, smapping.MapFields(&updateMahasiswaRequest))

	if err != nil {
		return nil, err
	}

	mahasiswa.User.ID = uid
	mahasiswa, err = c.mahasiswaRepo.UpdateMahasiswa(mahasiswa)

	if err != nil {
		return nil, err
	}

	res := _mahasiswa.NewMahasiswaResponse(mahasiswa)
	return &res, nil
}
