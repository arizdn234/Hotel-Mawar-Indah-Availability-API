package api

import (
	"net/http"
	"strconv"

	"github.com/arizdn234/hotel-mawar-indah-availability-api/internal/model"
	"github.com/arizdn234/hotel-mawar-indah-availability-api/internal/service"
	"github.com/gin-gonic/gin"
)

type RoomHandler struct {
	RoomService        *service.RoomService
	ReservationService *service.ReservationService
}

func NewRoomHandler(roomService *service.RoomService) *RoomHandler {
	return &RoomHandler{RoomService: roomService}
}

func (h *RoomHandler) GetAllRooms(c *gin.Context) {
	rooms, err := h.RoomService.GetAllRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rooms)
}

func (h *RoomHandler) GetRoomByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return
	}

	room, err := h.RoomService.GetRoomByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	c.JSON(http.StatusOK, room)
}

func (h *RoomHandler) CreateRoom(c *gin.Context) {
	var room model.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.RoomService.CreateRoom(&room); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Room created successfully"})
}

func (h *RoomHandler) UpdateRoom(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return
	}

	var room model.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	room.ID = uint(id)
	if err := h.RoomService.UpdateRoom(&room); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room updated successfully"})
}

func (h *RoomHandler) DeleteRoom(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return
	}

	if err := h.RoomService.DeleteRoom(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room deleted successfully"})
}

func (h *RoomHandler) RoomAvailabilityCheck(c *gin.Context) {
	availableRooms, err := h.RoomService.GetAvailableRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get available rooms"})
		return
	}

	if len(availableRooms) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No rooms available"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"available_rooms": availableRooms})
}

func (h *RoomHandler) RoomReservation(c *gin.Context) {
	var reservation model.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.RoomService.ReserveRoom(&reservation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make room reservation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room reservation successful", "reservation": reservation})
}

func (h *RoomHandler) ReservationCancellation(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return
	}

	if err := h.ReservationService.CancelReservation(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel reservation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation cancellation endpoint", "reservation_id": id})
}
