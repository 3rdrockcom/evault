package models

import "time"

// DataStore contains information about a stored data entry
type DataStore struct {
	ID          int       `json:"id"`
	Data        string    `json:"value"`
	Signature   string    `json:"-"`
	UserID      int       `json:"-"`
	PartitionID int       `json:"partition_id"`
	DateCreated time.Time `json:"-"`
}

// TableName gets the name of the database table
func (ds DataStore) TableName() string {
	return "datastore"
}
