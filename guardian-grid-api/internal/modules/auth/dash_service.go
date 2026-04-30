package auth

import (
	"errors"

	"github.com/pquerna/otp/totp"
)

type DService struct {
	repo *DRepository
}

func DashService(repo *DRepository) *DService {
	return &DService{repo: repo}
}

func (s *DService) Login(username, code string) (*User, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	// Validate TOTP
	if !totp.Validate(code, user.TOTPSecret) {
		return nil, errors.New("invalid TOTP code")
	}

	return user, nil
}
