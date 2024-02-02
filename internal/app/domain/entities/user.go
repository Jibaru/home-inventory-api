package entities

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"time"
)

var (
	ErrUserEmailExceedsMaxLengthOf100Chars                 = errors.New("email exceeds maximum length of 100 characters")
	ErrUserInvalidEmailAddress                             = errors.New("invalid email address")
	ErrUserPasswordMustBeBetween6And100Chars               = errors.New("password must be between 6 and 100 characters")
	ErrUserPasswordMustContainAtLeastOneLetterAndOneNumber = errors.New("password must contain at least one letter and one number")
)

type User struct {
	ID        string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(email string, password string) (*User, error) {
	if err := validateEmail(email); err != nil {
		return nil, err
	}

	if err := validatePassword(password); err != nil {
		return nil, err
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, errors.New("cannot hash password")
	}

	user := &User{
		ID:        uuid.NewString(),
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u *User) HasEqualPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func validateEmail(email string) error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if ok, _ := regexp.MatchString(emailRegex, email); !ok {
		return ErrUserInvalidEmailAddress
	}

	if len(email) > 100 {
		return ErrUserEmailExceedsMaxLengthOf100Chars
	}

	return nil
}

func validatePassword(password string) error {
	if len(password) < 6 || len(password) > 100 {
		return ErrUserPasswordMustBeBetween6And100Chars
	}

	hasLetter := false
	hasNumber := false
	for _, char := range password {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			hasLetter = true
		}
		if char >= '0' && char <= '9' {
			hasNumber = true
		}
	}

	if !hasLetter || !hasNumber {
		return ErrUserPasswordMustContainAtLeastOneLetterAndOneNumber
	}

	return nil
}
