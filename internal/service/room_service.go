package service

import (
	"github.com/arizdn234/hotel-mawar-indah-availability-api/internal/model"
	"github.com/arizdn234/hotel-mawar-indah-availability-api/internal/repository"
)

type RoomService struct {
	RoomRepository *repository.RoomRepository
}

func NewRoomService(roomRepository *repository.RoomRepository) *RoomService {
	return &RoomService{RoomRepository: roomRepository}
}

func (s *RoomService) GetAllRooms() ([]model.Room, error) {
	return s.RoomRepository.GetAllRooms()
}

func (s *RoomService) GetRoomByID(id uint) (*model.Room, error) {
	return s.RoomRepository.GetRoomByID(id)
}

func (s *RoomService) CreateRoom(room *model.Room) error {
	return s.RoomRepository.CreateRoom(room)
}

func (s *RoomService) UpdateRoom(room *model.Room) error {
	return s.RoomRepository.UpdateRoom(room)
}

func (s *RoomService) DeleteRoom(id uint) error {
	return s.RoomRepository.DeleteRoom(id)
}
