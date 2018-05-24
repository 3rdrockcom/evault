package datastore

import "errors"

var (
	// ErrEntryInvalid is an error given when the entry ID is not valid
	ErrEntryInvalid = errors.New("Invalid Entry ID")

	// ErrEntryNotFound is an error for a non-existent datastore entry
	ErrEntryNotFound = errors.New("Entry ID not found")

	// ErrEntryInvalidSignature is an error datastore entry does not match its stored signature
	ErrEntryInvalidSignature = errors.New("Cannot verify signature for this entry")
)
