package repository

import (
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
