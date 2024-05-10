package repository

import (
	"github.com/arizdn234/hotel-mawar-indah-availability-api/internal/model"
	"gorm.io/gorm"
)

type RoomRepository struct {
	DB *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{DB: db}
}

func (r *RoomRepository) GetAllRooms() ([]model.Room, error) {
	var rooms []model.Room

	if err := r.DB.Find(&rooms).Error; err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r *RoomRepository) GetRoomByID(id uint) (*model.Room, error) {
	var room model.Room

	if err := r.DB.First(&room, id).Error; err != nil {
		return nil, err
	}

	return &room, nil
}

func (r *RoomRepository) CreateRoom(room *model.Room) error {
	if err := r.DB.Create(room).Error; err != nil {
		return err
	}

	return nil
}

func (r *RoomRepository) UpdateRoom(room *model.Room) error {
	if err := r.DB.Save(room).Error; err != nil {
		return err
	}

	return nil
}

func (r *RoomRepository) DeleteRoom(id uint) error {
	if err := r.DB.Delete(&model.Room{}, id).Error; err != nil {
		return err
	}

	return nil
}
