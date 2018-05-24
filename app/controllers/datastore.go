package controllers

import (
	"net/http"
	"strconv"

	DataStore "github.com/epointpayment/evault/app/services/datastore"

	"github.com/labstack/echo"
)

// GetEntry retrieves a datastore entry from the database
func (co Controllers) GetEntry(c echo.Context) error {
	// Get entry ID
	entryID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		err = DataStore.ErrEntryInvalid
		return err
	}

	// Initialize datastore service
	userID := c.Get("userID").(int)
	partitionID := c.Get("partitionID").(int)
	dss := DataStore.New(userID, partitionID)

	// Get datastore entry by entry ID
	res, err := dss.Get(entryID)
	if err != nil {
		return err
	}

	// Send results
	return SendResponse(c, http.StatusOK, res)
}

// StoreEntry creates a new datastore entry and stores it in the database
func (co Controllers) StoreEntry(c echo.Context) error {
	userID := c.Get("userID").(int)
	partitionID := c.Get("partitionID").(int)

	value := c.FormValue("value")

	// Initialize datastore service
	dss := DataStore.New(userID, partitionID)

	// Store datastore entry
	res, err := dss.Store(value)
	if err != nil {
		return err
	}

	// Send results
	return SendResponse(c, http.StatusOK, res)
}
