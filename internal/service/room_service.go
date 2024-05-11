package service

import (
	"errors"

	"github.com/arizdn234/hotel-mawar-indah-availability-api/internal/model"
	"github.com/arizdn234/hotel-mawar-indah-availability-api/internal/repository"
)

type RoomService struct {
	RoomRepository        *repository.RoomRepository
	ReservationRepository *repository.ReservationRepository
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

func (s *RoomService) GetAvailableRooms() ([]model.Room, error) {
	availableRooms, err := s.RoomRepository.GetAvailableRooms()
	if err != nil {
		return nil, err
	}

	return availableRooms, nil
}

func (s *RoomService) ReserveRoom(reservation *model.Reservation) error {
	room, err := s.RoomRepository.GetRoomByID(reservation.RoomID)
	if err != nil {
		return err
	}

	if !room.Available {
		return errors.New("room is not available")
	}

	room.Available = false
	if err := s.RoomRepository.UpdateRoom(room); err != nil {
		return err
	}

	if err := s.ReservationRepository.CreateReservation(reservation); err != nil {
		room.Available = true
		if err := s.RoomRepository.UpdateRoom(room); err != nil {
			return err
		}
		return err
	}

	return nil
}
