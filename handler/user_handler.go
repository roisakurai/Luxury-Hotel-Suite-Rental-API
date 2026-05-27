package handler

import (
	"net/http"
	"p2-ip-hotel-rental/helper"
	"p2-ip-hotel-rental/models"
	"p2-ip-hotel-rental/service"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

// Register godoc
// @Summary Register new user
// @Description Create new user account
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.RegisterRequest true "Register Request"
// @Success 200 {object} models.MessageResponse
// @Failure 400 {object} map[string]string
// @Router /register [post]
func (h *UserHandler) Register(c echo.Context) error {
	var req models.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request",
		})
	}

	err := h.userService.Register(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "register success",
	})
}

// Login godoc
// @Summary Login user
// @Description Login and get JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.LoginRequest true "Login Request"
// @Success 200 {object} models.LoginResponse
// @Failure 401 {object} map[string]string
// @Router /login [post]
func (h *UserHandler) Login(c echo.Context) error {
	var req models.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request",
		})
	}

	user, err := h.userService.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
	}

	token, err := helper.GenerateToken(user.ID)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "failed to generate token"})
	}

	userRes := models.ToUserResponse(user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "login success",
		"data":    userRes,
		"token":   token,
	})
}

// TopUp godoc
// @Summary Top up user balance
// @Description Add balance to authenticated user
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body models.TopUpRequest true "Top Up Request"
// @Success 200 {object} models.MessageResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /top-up [post]
func (h *UserHandler) TopUp(c echo.Context) error {
	var req models.TopUpRequest

	userIDInterface := c.Get("user_id")
	userID, ok := userIDInterface.(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "invalid user",
		})
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request",
		})
	}

	err := h.userService.TopUp(userID, req.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "top up success",
	})
}

// GetProfile godoc
// @Summary Get user profile
// @Description Get authenticated user profile including deposit amount
// @Tags User
// @Security BearerAuth
// @Produce json
// @Success 200 {object} models.UserProfileWrapper
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /profile [get]
func (h *UserHandler) GetProfile(c echo.Context) error {

	userIDInterface := c.Get("user_id")
	userID, ok := userIDInterface.(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "invalid user",
		})
	}

	user, err := h.userService.GetProfile(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "user not found",
		})
	}

	res := models.ToUserProfileResponse(user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    res,
	})
}
