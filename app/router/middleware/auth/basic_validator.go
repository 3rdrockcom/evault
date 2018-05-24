package auth

import (
	"strconv"

	API "github.com/epointpayment/evault/app/services/api"
	User "github.com/epointpayment/evault/app/services/user"

	"github.com/labstack/echo"
)

// BasicValidator is a validator used for basic auth middleware
func BasicValidator(username, password string, c echo.Context) (isValid bool, err error) {
	// Initialize API service
	sa := API.New()

	// Get program ID
	programID, err := strconv.Atoi(c.QueryParam("program_id"))
	if err != nil {
		err = User.ErrInvalidProgramID
		return
	}

	// Check is user is authorized
	isValid, err = sa.DoAuth(username, password, programID)
	if !isValid || err != nil {
		return
	}

	// Initialize API service
	us := User.New()

	// User is authorized, get user information
	user, err := us.GetByUsername(username, programID)
	if err != nil {
		return
	}

	// Pass user information to context
	c.Set("userID", user.ID)
	return
}
