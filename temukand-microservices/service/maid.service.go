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

	_maid "tmi-gin/service/maid"
)

type MaidService interface {
	All(userID string) (*[]_maid.MaidResponse, error)
	CreateMaid(maidRequest dto.CreateMaidRequest, userID string) (*_maid.MaidResponse, error)
	UpdateMaid(updateMaidRequest dto.UpdateMaidRequest, userID string) (*_maid.MaidResponse, error)
	FindOneMaidByID(maidID string) (*_maid.MaidResponse, error)
	DeleteMaid(maidID string, userID string) error
}

type maidService struct {
	maidRepo repo.MaidRepository
}

func NewMaidService(maidRepo repo.MaidRepository) MaidService {
	return &maidService{
		maidRepo: maidRepo,
	}
}

func (c *maidService) All(userID string) (*[]_maid.MaidResponse, error) {
	maids, err := c.maidRepo.All(userID)
	if err != nil {
		return nil, err
	}

	prods := _maid.NewMaidArrayResponse(maids)
	return &prods, nil
}

func (c *maidService) CreateMaid(maidRequest dto.CreateMaidRequest, userID string) (*_maid.MaidResponse, error) {
	maid := entity.Maid{}
	err := smapping.FillStruct(&maid, smapping.MapFields(&maidRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	id, _ := strconv.ParseInt(userID, 0, 64)
	maid.UserID = id
	p, err := c.maidRepo.InsertMaid(maid)
	if err != nil {
		return nil, err
	}

	res := _maid.NewMaidResponse(p)
	return &res, nil
}

func (c *maidService) FindOneMaidByID(maidID string) (*_maid.MaidResponse, error) {
	maid, err := c.maidRepo.FindOneMaidByID(maidID)

	if err != nil {
		return nil, err
	}

	res := _maid.NewMaidResponse(maid)
	return &res, nil
}

func (c *maidService) UpdateMaid(updateMaidRequest dto.UpdateMaidRequest, userID string) (*_maid.MaidResponse, error) {
	maid, err := c.maidRepo.FindOneMaidByID(fmt.Sprintf("%d", updateMaidRequest.ID))
	if err != nil {
		return nil, err
	}

	uid, _ := strconv.ParseInt(userID, 0, 64)
	if maid.UserID != uid {
		return nil, errors.New("Ini bukan milik anda")
	}

	maid = entity.Maid{}
	err = smapping.FillStruct(&maid, smapping.MapFields(&updateMaidRequest))

	if err != nil {
		return nil, err
	}

	maid.UserID = uid
	maid, err = c.maidRepo.UpdateMaid(maid)

	if err != nil {
		return nil, err
	}

	res := _maid.NewMaidResponse(maid)
	return &res, nil
}

func (c *maidService) DeleteMaid(maidID string, userID string) error {
	maid, err := c.maidRepo.FindOneMaidByID(maidID)
	if err != nil {
		return err
	}

	if fmt.Sprintf("%d", maid.UserID) != userID {
		return errors.New("Ini bukan milik anda")
	}

	c.maidRepo.DeleteMaid(maidID)
	return nil

}
