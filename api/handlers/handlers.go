package handlers

import (
	Log "backend/pkg/helpers/log"
	"fmt"
	"net/http"
)

// ShowLocations shows active locations from the database.
// @Summary Show active locations
// @Description Retrieve a list of active locations
// @Tags locations
// @ID ShowLocations
// @Produce json
// @Success 200 {object} []models.Locations
// @Failure 500 {string} string "failed to find data"
// @Failure 500 {string} string "an error occurred while retrieving data from the database"
// @Failure 500 {string} string "no data found"
// @Router /api/locations/show [get]
func (s *Server) ShowLocations(w http.ResponseWriter, r *http.Request) {
	result, err := s.DB.Mongo.ShowLocations()
	if Log.Err(err) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Render(w, 200, result)
}

// ShowReservedCars retrieves the list of reserved cars from the database.
// @Summary Show reserved cars
// @Description Retrieve the list of reserved cars from the database.
// @Tags cars
// @ID ShowReservedCars
// @Produce json
// @Success 200 {object} []models.Cars
// @Failure 500 {string} string "failed to find data"
// @Failure 500 {string} string "an error occurred while retrieving data from the database"
// @Failure 500 {string} string "no data found"
// @Router /api/cars/showreservedcars [get]
func (s *Server) ShowReservedCars(w http.ResponseWriter, r *http.Request) {
	result, err := s.DB.Mongo.ShowReservedCars()
	if Log.Err(err) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Render(w, 200, result)
}

// ReserveCar is a handler function for reserving a car by ID.
// @Summary Reserve a car
// @Description Reserve a car by ID and user information
// @Tags cars
// @Produce json
// @Param id query string true "Car ID"
// @Param user body models.User true "User information"
// @Success 201 {string} string "successfully reserved car {id}"
// @Failure 400 {string} string "missing required parameters"
// @Failure 400 {string} string "error occurred while converting data"
// @Failure 400 {string} string "empty request body"
// @Failure 400 {string} string "error occurred while decoding JSON"
// @Failure 500 {string} string "failed to update data"
// @Failure 500 {string} string "the car is already reserved"
// @Router /api/cars/reserve [post]
func (s *Server) ReserveCar(w http.ResponseWriter, r *http.Request) {
	id, err := ParseQuery_ReserveCar(r)
	if Log.Err(err) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := ReadJSON(r)
	if Log.Err(err) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := s.DB.Mongo.ReserveCar(id, user); Log.Err(err) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render := fmt.Sprintf("successfully reserved car %d", id)
	Render(w, 201, render)
}

// AvailableCars returns a list of available cars for a given date, time(hour), and location.
// @Summary Available Cars
// @Description Get a list of available cars by options(param)
// @Tags cars
// @Produce json
// @Param receiver_date query string true "Date in the format of yyyy-mm-dd"
// @Param delivery_date query string true "Date in the format of yyyy-mm-dd"
// @Param time_start(hour) query string true "Time in the format of hh.mm"
// @Param time_end(hour) query string true "Time in the format of hh.mm"
// @Param location query string true "Location by id"
// @Success 200 {object} []models.Cars
// @Failure 400 {string} string "missing required parameters"
// @Failure 400 {string} string "error occurred while converting data"
// @Failure 500 {string} string "failed to find data"
// @Failure 500 {string} string "an error occurred while retrieving data from the database"
// @Failure 500 {string} string "no active offices available"
// @Failure 500 {string} string "no data found"
// @Router /api/cars/available [get]
func (s *Server) AvailableCars(w http.ResponseWriter, r *http.Request) {
	data, err := ParseQuery_AvailableCars(r)
	if Log.Err(err) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := s.DB.Mongo.AvailableCars(data)
	if Log.Err(err) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Render(w, 200, result)
}
