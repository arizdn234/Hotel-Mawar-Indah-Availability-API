package service

import "github.com/arizdn234/hotel-mawar-indah-availability-api/internal/repository"

type ReservationService struct {
	ReservationRepository *repository.ReservationRepository
}

func NewReservationService(reservationRepository *repository.ReservationRepository) *ReservationService {
	return &ReservationService{ReservationRepository: reservationRepository}
}

func (s *ReservationService) CancelReservation(id uint64) error {
	err := s.ReservationRepository.CancelReservation(uint(id))
	if err != nil {
		return err
	}
	return nil
}
