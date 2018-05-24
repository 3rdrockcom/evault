package router

import (
	"net/http"

	"github.com/epointpayment/evault/app/controllers"
	API "github.com/epointpayment/evault/app/services/api"
	DataStore "github.com/epointpayment/evault/app/services/datastore"
	User "github.com/epointpayment/evault/app/services/user"

	"github.com/labstack/echo"
)

// appendErrorHandler handles errors for the router
func (r *Router) appendErrorHandler() {
	r.e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		message := err.Error()
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			message = he.Message.(string)
		}

		// Override status code based on error responses
		switch message {
		case API.ErrInvalidCredentials.Error():
			code = http.StatusForbidden
		case DataStore.ErrInvalidPartitionID.Error():
			code = http.StatusBadRequest
		case DataStore.ErrEntryNotFound.Error():
			code = http.StatusNotFound
		case DataStore.ErrEntryInvalid.Error():
			code = http.StatusBadRequest
		case DataStore.ErrEntryInvalidSignature.Error():
			// Status Code 419: Checksum Failed
			code = 419
		case User.ErrUserNotFound.Error():
			code = http.StatusNotFound
		case User.ErrInvalidProgramID.Error():
			code = http.StatusBadRequest
		case User.ErrUserExists.Error():
			code = http.StatusBadRequest
		default:
			// Unknown error
			if _, ok := err.(*echo.HTTPError); !ok {
				message = "Internal Error"
			}
		}

		// Send error in a specific format
		controllers.SendErrorResponse(c, code, message)

		// Log errors
		c.Logger().Error(err)
	}
}
