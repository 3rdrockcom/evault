package user

import (
	"database/sql"
	"time"

	"github.com/epointpayment/evault/app/models"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/go-sql-driver/mysql"
)

// DB is the database handler
var DB *dbx.DB

// UserService is a service that manages a user
type UserService struct {
	UserID int
}

// New creates an instance of the service
func New() *UserService {

	return &UserService{}
}

// Load creates an instance of the service for a particular user
func Load(userID int) *UserService {
	return &UserService{
		UserID: userID,
	}
}

// Create creates a user
func (us *UserService) Create(username string, password string, programID int) (user *models.User, err error) {
	user = new(models.User)

	if programID <= 0 {
		err = ErrInvalidProgramID
		return
	}

	// Insert into database
	tx, err := DB.Begin()
	if err != nil {
		return
	}

	user.Username = username
	user.Password = password
	user.ProgramID = programID
	user.DateCreated = time.Now().UTC()

	err = tx.Model(user).Insert()
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			err = ErrUserExists
		}
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}

// Get gets a user by id
func (us *UserService) Get(userID int, programID int) (user *models.User, err error) {
	user = new(models.User)

	err = DB.Select().
		Where(dbx.HashExp{"id": userID}).
		One(user)
	if err != nil {
		if err == sql.ErrNoRows {
			err = ErrUserNotFound
		}
		return nil, err
	}

	return
}

// GetByUsername gets a user by username
func (us *UserService) GetByUsername(username string, programID int) (user *models.User, err error) {
	user = new(models.User)

	err = DB.Select().
		Where(dbx.HashExp{"username": username, "program_id": programID}).
		One(user)
	if err != nil {
		if err == sql.ErrNoRows {
			err = ErrUserNotFound
		}
		return nil, err
	}

	return
}
