package models

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TopUpRequest struct {
	Amount float64 `json:"amount"`
}

type BookingRequest struct {
	SuiteID  uint   `json:"suite_id"`
	CheckIn  string `json:"check_in"`
	CheckOut string `json:"check_out"`
}
