**Project Name: Hotel Mawar Indah Availability API**

**Description:**
Hotel Mawar Indah Availability API is a backend service that provides information about hotel room availability. This service is designed to be used by booking applications or booking service providers from various hotels, as well as internal company applications. With this API, clients can check the availability of rooms that are not currently booked and make reservations.

**Features:**
1. **Room Availability Check:** Clients can check the availability of hotel rooms that are currently unbooked.
2. **Room Reservation:** Clients can make reservations for available hotel rooms.
3. **Reservation Cancellation:** Clients can cancel reservations that have been made.
4. **Room Information:** Clients can obtain detailed information about hotel rooms, such as facilities, prices, and more.

**Tech Stack:**
- **Programming Language:** Go (Golang) v1.22
- **Web Framework:** Gin
- **ORM:** GORM
- **Database:** PostgreSQL

**Installation:**
1. Make sure Go is installed on your computer. To install Go, follow the official guide at [https://golang.org/doc/install](https://golang.org/doc/install).
2. Clone this repository to your local directory:
   ```
   git clone <repository-link>
   ```
3. Open a terminal and navigate to the repository directory:
   ```
   cd hotel-mawar-indah-api
   ```
4. Install dependencies using Go Module:
   ```
   go mod tidy
   ```
5. Create the `.env` file from the `.env.example` file and adjust your database configuration:
   ```
   cp .env.example .env
   ```
6. Run the application:
   ```
   go run main.go
   ```

**Usage:**
1. **Room Availability Check:**
   - Endpoint: `/availability`
   - Method: `GET`
   - Description: Checks the availability of hotel rooms based on the provided parameters.

2. **Room Reservation:**
   - Endpoint: `/reservation`
   - Method: `POST`
   - Description: Makes a reservation for a hotel room by providing booking data.

3. **Reservation Cancellation:**
   - Endpoint: `/reservation/:id/cancel`
   - Method: `DELETE`
   - Description: Cancels a hotel room reservation based on the reservation ID.

4. **Room Information:**
   - Endpoint: `/room/:id`
   - Method: `GET`
   - Description: Retrieves detailed information about a hotel room based on the room ID.

**Contribution:**
If you would like to contribute to this project, please open a new issue to discuss the changes you would like to make. You can also submit a pull request with proposed changes.

**License:**
This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

**Contact:**
If you have any questions or feedback, feel free to contact the project team via email: contact@hotelmawarindah.com.

**Thank you for your interest and support for the Hotel Mawar Indah Availability API!**