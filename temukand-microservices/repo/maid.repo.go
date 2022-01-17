package repo

import (
	"tmi-gin/entity"

	"gorm.io/gorm"
)

type MaidRepository interface {
	All(userID string) ([]entity.Maid, error)
	InsertMaid(maid entity.Maid) (entity.Maid, error)
	UpdateMaid(maid entity.Maid) (entity.Maid, error)
	DeleteMaid(maidID string) error
	FindOneMaidByID(ID string) (entity.Maid, error)
	FindAllMaid(userID string) ([]entity.Maid, error)
}

type maidRepo struct {
	connection *gorm.DB
}

func NewMaidRepo(connection *gorm.DB) MaidRepository {
	return &maidRepo{
		connection: connection,
	}
}

func (c *maidRepo) All(userID string) ([]entity.Maid, error) {
	maids := []entity.Maid{}
	c.connection.Preload("User").Where("user_id = ?", userID).Find(&maids)
	return maids, nil
}

func (c *maidRepo) InsertMaid(maid entity.Maid) (entity.Maid, error) {
	c.connection.Save(&maid)
	c.connection.Preload("User").Find(&maid)
	return maid, nil
}

func (c *maidRepo) UpdateMaid(maid entity.Maid) (entity.Maid, error) {
	c.connection.Save(&maid)
	c.connection.Preload("User").Find(&maid)
	return maid, nil
}

func (c *maidRepo) FindOneMaidByID(maidID string) (entity.Maid, error) {
	var maid entity.Maid
	res := c.connection.Preload("User").Where("id = ?", maidID).Take(&maid)
	if res.Error != nil {
		return maid, res.Error
	}
	return maid, nil
}

func (c *maidRepo) FindAllMaid(userID string) ([]entity.Maid, error) {
	maids := []entity.Maid{}
	c.connection.Where("user_id = ?", userID).Find(&maids)
	return maids, nil
}

func (c *maidRepo) DeleteMaid(maidID string) error {
	var maid entity.Maid
	res := c.connection.Preload("User").Where("id = ?", maidID).Take(&maid)
	if res.Error != nil {
		return res.Error
	}
	c.connection.Delete(&maid)
	return nil
}
