[![Open in Visual Studio Code](https://classroom.github.com/assets/open-in-vscode-2e0aaae1b6195c2367325f4f02e2d04e9abb55f0b24a779b69b11b9e10269abc.svg)](https://classroom.github.com/online_ide?assignment_repo_id=23771909&assignment_repo_type=AssignmentRepo)
# Luxury Hotel Suite Rental API
REST API project for managing luxury hotel suite rentals with JWT authentication, booking system, balance top-up, booking reports, and email notifications using SendGrid.

## Features
User Registration
User Login with JWT Authentication
Top Up Balance
Get User Profile
Get Available Suites
Book Luxury Hotel Suites
Booking Report / History
Email Notification Integration (SendGrid)
Swagger API Documentation
Unit Testing with Mocking
Railway Deployment Ready

## Tech Stack
Golang
Echo Framework
GORM
PostgreSQL
JWT Authentication
Swagger / OpenAPI
SendGrid API
Testify Mock
Railway

## Project Structure

```txt
project/
├── config/
├── docs/
├── handler/
├── helper/
├── middleware/
├── mocks/
├── models/
├── repository/
├── routes/
├── service/
├── main.go
├── .env
└── README.md
```


## Installation
1. Clone Repository
git clone https://github.com/roisakurai/Luxury-Hotel-Suite-Rental-API
cd Luxury-Hotel-Suite-Rental-API

2. Install Dependencies
go mod tidy

3. Setup Environment Variables
Create .env file:
DATABASE_URL=your_postgresql_url
JWT_SECRET=your_jwt_secret
PORT=8080

SENDGRID_API_KEY=your_sendgrid_api_key
EMAIL_SENDER=your_verified_email

## Database Migration
Run PostgreSQL and ensure database is connected.

Then run application:
go run main.go

## API Endpoints
### Public Endpoints
Method	Endpoint	Description
POST	/register	Register user
POST	/login	Login user
GET	/suites	Get all suites

### Protected Endpoints (JWT Required)
Method	Endpoint	Description
POST	/top-up	Top up user balance
POST	/bookings	Create booking
GET	/booking-report	Get booking history
GET	/profile	Get user profile
Authentication

## Use JWT token from login response.
Example:
Authorization: Bearer your_token

## Example Request
### Register
POST /register

```txt
{
  "email": "user@mail.com",
  "password": "123456"
}
```

### Login
POST /login

```txt
{
  "email": "user@mail.com",
  "password": "123456"
}
```

### Top Up
POST /top-up

```txt
{
  "amount": 1000000
}
```

### Create Booking
POST /bookings

```txt
{
  "suite_id": 1,
  "check_in": "2026-05-01",
  "check_out": "2026-05-03"
}
```

## Swagger Documentation
Generate Swagger docs:
swag init

Run application:
go run main.go

Open Swagger UI:
http://localhost:8080/swagger/index.html

## Unit Testing
Run all tests:

```txt
go test ./... -v
```

## Email Notification
This project integrates with SendGrid API for:
Registration success email
Booking confirmation email