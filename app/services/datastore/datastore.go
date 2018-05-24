package datastore

import (
	"database/sql"
	"time"

	"github.com/epointpayment/evault/app/models"
	"github.com/epointpayment/evault/app/services/signature"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// DB is the database handler
var DB *dbx.DB

// DataStoreService is a service that manages the datastore
type DataStoreService struct {
	userID    int
	signature *signature.SignatureService
}

// New creates an instance of the service
func New(userID int) *DataStoreService {
	ss := signature.New()

	return &DataStoreService{
		userID:    userID,
		signature: ss,
	}
}

// Store creates a new datastore entry and stores it in the database
func (dss *DataStoreService) Store(str string) (datastore *models.DataStore, err error) {
	datastore = new(models.DataStore)

	// Generate a hash
	hash, err := dss.signature.Create(str)
	if err != nil {
		return
	}

	// Insert into database
	tx, err := DB.Begin()
	if err != nil {
		return
	}

	datastore.UserID = dss.userID
	datastore.Data = str
	datastore.Signature = hash
	datastore.DateCreated = time.Now().UTC()

	err = tx.Model(datastore).Insert()
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}

// Get retrieves a datastore entry from the database
func (dss *DataStoreService) Get(entryID int) (datastore *models.DataStore, err error) {
	datastore = new(models.DataStore)

	err = DB.Select().
		Where(dbx.HashExp{"id": entryID, "user_id": dss.userID}).
		One(datastore)
	if err != nil {
		if err == sql.ErrNoRows {
			err = ErrEntryNotFound
		}
		return nil, err
	}

	isValid, err := dss.signature.Verify(datastore.Data, datastore.Signature)
	if err != nil && !isValid {
		err = ErrEntryInvalidSignature
		return nil, err
	}

	return
}
