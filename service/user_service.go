package service

import (
	"errors"
	"p2-ip-hotel-rental/models"
	"p2-ip-hotel-rental/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(email, password string) error
	Login(email, password string) (*models.User, error)
	TopUp(userID uint, amount float64) error
	GetProfile(userID uint) (*models.User, error)
}

type userService struct {
	userRepo repository.UserRepository
	txRepo   repository.TransactionRepository
	emailSvc EmailService
}

func NewUserService(
	userRepo repository.UserRepository,
	txRepo repository.TransactionRepository,
	emailSvc EmailService,
) UserService {
	return &userService{
		userRepo: userRepo,
		txRepo:   txRepo,
		emailSvc: emailSvc,
	}
}

func (s *userService) Register(email, password string) error {

	_, err := s.userRepo.FindByEmail(email)
	if err == nil {
		return errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	user := &models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return err
	}

	// send email
	if s.emailSvc != nil {
		go func() {
			_ = s.emailSvc.SendEmail(
				email,
				"Welcome to Sakurai Luxury Hotel Rental",
				"Your registration is successful",
			)
		}()
	}

	return nil
}

func (s *userService) Login(email, password string) (*models.User, error) {

	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

func (s *userService) TopUp(userID uint, amount float64) error {

	if amount <= 0 {
		return errors.New("invalid amount")
	}

	err := s.userRepo.UpdateDeposit(userID, amount)
	if err != nil {
		return err
	}

	tx := &models.Transaction{
		UserID: userID,
		Amount: amount,
		Type:   "topup",
	}

	return s.txRepo.Create(tx)
}

func (s *userService) GetProfile(userID uint) (*models.User, error) {
	return s.userRepo.FindByID(userID)
}
