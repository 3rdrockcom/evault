package datastore

import "errors"

var (
	// ErrInvalidPartitionID is an error given when the partition ID is not valid
	ErrInvalidPartitionID = errors.New("Invalid Partition ID")

	// ErrEntryInvalid is an error given when the entry ID is not valid
	ErrEntryInvalid = errors.New("Invalid Entry ID")

	// ErrEntryNotFound is an error for a non-existent datastore entry
	ErrEntryNotFound = errors.New("Entry ID not found")

	// ErrEntryInvalidSignature is an error given when a datastore entry does not match its stored signature
	ErrEntryInvalidSignature = errors.New("Cannot verify signature for this entry")

	// ErrEntryLengthExceeded is an error given when a datastore entry exceeds a specified length
	ErrEntryLengthExceeded = errors.New("Entry value is too long")
)
