package main

import (
	"log"
	"os"
	"p2-ip-hotel-rental/config"
	"p2-ip-hotel-rental/handler"
	"p2-ip-hotel-rental/middleware"
	"p2-ip-hotel-rental/repository"
	"p2-ip-hotel-rental/service"

	_ "p2-ip-hotel-rental/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Luxury Hotel Suite Rental
// @description API for Luxury Hotel Suite Rental
// @host p2-ip-roisakurai-production.up.railway.app
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, sqlDB := config.ConnectDB()
	defer sqlDB.Close()

	e := echo.New()

	// test api
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "API is running",
		})
	})

	// test db
	e.GET("/test-db", func(c echo.Context) error {
		row := sqlDB.QueryRow("SELECT 1")

		var result int
		err := row.Scan(&result)
		if err != nil {
			return c.JSON(500, map[string]string{
				"message": "DB error",
			})
		}

		return c.JSON(200, map[string]interface{}{
			"message": "DB connected",
			"result":  result,
		})
	})

	userRepo := repository.NewUserRepository(db)
	suiteRepo := repository.NewSuiteRepository(db)
	bookingRepo := repository.NewBookingRepository(db)
	txRepo := repository.NewTransactionRepository(db)
	emailService := service.NewEmailService()

	_ = userRepo
	_ = suiteRepo
	_ = bookingRepo
	_ = txRepo
	_ = emailService

	userService := service.NewUserService(userRepo, txRepo, emailService)

	suiteService := service.NewSuiteService(suiteRepo)

	bookingService := service.NewBookingService(
		db,
		userRepo,
		suiteRepo,
		bookingRepo,
		txRepo,
		emailService,
	)

	_ = userService
	_ = suiteService
	_ = bookingService

	userHandler := handler.NewUserHandler(userService)
	bookingHandler := handler.NewBookingHandler(bookingService)
	suiteHandler := handler.NewSuiteHandler(suiteService)

	// public routes
	e.POST("/register", userHandler.Register)
	e.POST("/login", userHandler.Login)
	e.GET("/suites", suiteHandler.GetSuites)

	auth := e.Group("")
	auth.Use(middleware.JWTMiddleware)

	// protected routes
	auth.POST("/top-up", userHandler.TopUp)
	auth.POST("/bookings", bookingHandler.CreateBooking)
	auth.GET("/booking-report", bookingHandler.GetBookingReport)
	auth.GET("/profile", userHandler.GetProfile)

	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)

	e.Logger.Fatal(e.Start(":" + port))
}
