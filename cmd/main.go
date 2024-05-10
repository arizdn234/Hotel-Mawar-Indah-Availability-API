package main

import (
	"log"
	"net/http"
	"os"

	"github.com/arizdn234/hotel-mawar-indah-availability-api/db"
	"github.com/arizdn234/hotel-mawar-indah-availability-api/internal/api"
	"github.com/arizdn234/hotel-mawar-indah-availability-api/internal/repository"
	"github.com/arizdn234/hotel-mawar-indah-availability-api/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

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
		// make sure this endpoint is only available to the admin role
		// r.Use()

		rooms.GET("/", roomHandler.GetAllRooms)
		rooms.GET("/:id", roomHandler.GetRoomByID)
		rooms.POST("/", roomHandler.CreateRoom)
		rooms.PUT("/:id", roomHandler.UpdateRoom)
		rooms.DELETE("/:id", roomHandler.DeleteRoom)
	}

	r.GET("/availability", roomHandler.RoomAvailabilityCheck)
	r.POST("/reservation", roomHandler.RoomReservation)
	r.DELETE("/reservation/:id/cancel", roomHandler.ReservationCancellation)

	port := os.Getenv("PORT")

	log.Printf("Server is running on port %v\n\n`http://localhost:%v`", port, port)
	if err := r.Run(":" + port); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}
