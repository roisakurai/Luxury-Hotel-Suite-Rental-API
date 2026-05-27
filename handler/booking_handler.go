package handler

import (
	"net/http"
	"p2-ip-hotel-rental/models"
	"p2-ip-hotel-rental/service"
	"time"

	"github.com/labstack/echo/v4"
)

type BookingHandler struct {
	bookingService service.BookingService
}

func NewBookingHandler(bookingService service.BookingService) *BookingHandler {
	return &BookingHandler{bookingService}
}

// CreateBooking godoc
// @Summary Create booking
// @Tags Booking
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body models.BookingRequest true "Booking Request"
// @Success 200 {object} models.BookingResponse
// @Failure 400 {object} map[string]string
// @Router /bookings [post]
func (h *BookingHandler) CreateBooking(c echo.Context) error {
	var req models.BookingRequest
	userID := c.Get("user_id").(uint)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request",
		})
	}

	checkIn, err := time.Parse("2006-01-02", req.CheckIn)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid check_in format",
		})
	}

	checkOut, err := time.Parse("2006-01-02", req.CheckOut)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid check_out format",
		})
	}

	booking, err := h.bookingService.CreateBooking(
		userID,
		req.SuiteID,
		checkIn,
		checkOut,
	)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	bookingRes := models.ToBookingResponse(booking)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "booking success",
		"data":    bookingRes,
	})
}

// GetBookingReport godoc
// @Summary Get booking history
// @Tags Booking
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.BookingResponse
// @Failure 400 {object} map[string]string
// @Router /booking-report [get]
func (h *BookingHandler) GetBookingReport(c echo.Context) error {

	userIDInterface := c.Get("user_id")
	userID, ok := userIDInterface.(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "invalid user",
		})
	}

	bookings, err := h.bookingService.GetBookingReport(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	var result []models.BookingResponse
	for _, b := range bookings {
		result = append(result, models.ToBookingResponse(&b))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    result,
	})
}
