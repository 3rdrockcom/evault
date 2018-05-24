package api

import (
	User "github.com/epointpayment/evault/app/services/user"
)

// APIService is a service that manages the API access
type APIService struct{}

// New creates an instance of the service
func New() *APIService {
	return &APIService{}
}

// DoAuth checks if a user's auth credentials are valid
func (as *APIService) DoAuth(username, password string, programID int) (isValid bool, err error) {
	us := User.New()
	user, err := us.GetByUsername(username, programID)
	if err != nil {
		if err == User.ErrUserNotFound {
			err = ErrInvalidCredentials
		}
		return
	}

	if user.Password == password {
		isValid = true
	}

	return
}
