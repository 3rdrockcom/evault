package services

import (
	"github.com/epointpayment/evault/app/database"
	"github.com/epointpayment/evault/app/services/datastore"
	"github.com/epointpayment/evault/app/services/user"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// db is the database handler
var db *dbx.DB

// Services boots application-specific services
type Services struct{}

// New starts the service setup process
func New(DB *database.Database) error {
	db = DB.GetInstance()

	// Attach the database handler to service
	datastore.DB = db
	user.DB = db

	return nil
}
