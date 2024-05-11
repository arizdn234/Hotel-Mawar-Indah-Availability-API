package repository

import (
	"errors"

	"github.com/arizdn234/hotel-mawar-indah-availability-api/internal/model"
	"gorm.io/gorm"
)

type ReservationRepository struct {
	DB *gorm.DB
}

func NewReservationRepository(db *gorm.DB) *ReservationRepository {
	return &ReservationRepository{
		DB: db,
	}
}

func (r *ReservationRepository) CreateReservation(reservation *model.Reservation) error {
	if err := r.DB.Create(reservation).Error; err != nil {
		return err
	}
	return nil
}

func (r *ReservationRepository) CancelReservation(id uint) error {
	var reservation model.Reservation
	if err := r.DB.First(&reservation, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("reservation not found")
		}
		return err
	}
	reservation.Status = "canceled"
	if err := r.DB.Save(&reservation).Error; err != nil {
		return err
	}

	if err := r.Available(reservation.RoomID, true); err != nil {
		return err
	}

	return nil
}

func (r *ReservationRepository) Available(roomID uint, available bool) error {
	var room model.Room
	if err := r.DB.First(&room, roomID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("room not found")
		}
		return err
	}

	room.Available = available
	if err := r.DB.Save(&room).Error; err != nil {
		return err
	}

	return nil
}
