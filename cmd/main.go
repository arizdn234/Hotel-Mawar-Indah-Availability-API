package main

import (
	"log"

	"github.com/arizdn234/hotel-mawar-indah-availability-api/db"
	"github.com/arizdn234/hotel-mawar-indah-availability-api/internal/api"
	"github.com/arizdn234/hotel-mawar-indah-availability-api/internal/repository"
	"github.com/arizdn234/hotel-mawar-indah-availability-api/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	roomRepository := repository.NewRoomRepository(db.DB)
	roomService := service.NewRoomService(roomRepository)
	roomHandler := api.NewRoomHandler(roomService)

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	rooms := r.Group("/rooms")
	{
		// r.Use()

		rooms.GET("/", roomHandler.GetAllRooms)
		rooms.GET("/:id", roomHandler.GetRoomByID)
		rooms.POST("/", roomHandler.CreateRoom)
		rooms.PUT("/:id", roomHandler.UpdateRoom)
		rooms.DELETE("/:id", roomHandler.DeleteRoom)
	}

	// r.GET("/availability", RoomAvailabilityCheck)
	// r.POST("/reservation", RoomReservation)
	// r.DELETE("/reservation/:id/cancel", ReservationCancellation)

	r.Run(":8080")
}
