package handlers

import (
	"backend/models"
	"backend/pkg/helpers"
	Log "backend/pkg/helpers/log"
	"encoding/json"
	"net/http"
	"strconv"
)

// Render helper function to render page.
func Render(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// ReadJSON helper function to read JSON data.
func ReadJSON(r *http.Request) (data *models.User, err error) {
	if r.Body == nil {
		err = ErrEmptyRequestBody
		return
	}
	j := json.NewDecoder(r.Body)
	j.DisallowUnknownFields()
	if err = j.Decode(&data); Log.Err(err) {
		err = ErrDecodeJSON
		return
	}
	if helpers.IsParamNull(data.Name, data.SurName, data.Nation_Id, data.Phone_Number) {
		err = ErrMissingRequiredParams
		return
	}
	return
}

// ParseQuery_ReserveCar parses the query parameters from the request.
func ParseQuery_ReserveCar(r *http.Request) (id int, err error) {
	id_str := r.URL.Query().Get("id")
	if helpers.IsParamNull(id_str) {
		err = ErrMissingRequiredParams
		return
	}
	id, err = strconv.Atoi(id_str)
	if Log.Err(err) {
		err = ErrConvert
		return
	}
	return
}

// ParseQuery_AvailableCars parses the query parameters from the request.
func ParseQuery_AvailableCars(r *http.Request) (data *models.AvailableCarsParams, err error) {
	receiving_date := r.URL.Query().Get("receiver_date")
	delivery_date := r.URL.Query().Get("delivery_date")
	location := r.URL.Query().Get("location")
	data = new(models.AvailableCarsParams)
	data.TimeStart = r.URL.Query().Get("time_start")
	data.TimeEnd = r.URL.Query().Get("time_end")
	if helpers.IsParamNull(receiving_date, delivery_date, data.TimeStart, data.TimeEnd, location) {
		err = ErrMissingRequiredParams
		return
	}
	// convert date to int(weekday) and location to int
	if data.DateInt, err = helpers.DateToInt(receiving_date, delivery_date); Log.Err(err) {
		err = ErrConvert
		return
	}
	if data.Location, err = strconv.Atoi(location); Log.Err(err) {
		err = ErrConvert
		return
	}
	return
}
