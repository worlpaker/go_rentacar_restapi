package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	log.SetOutput(io.Discard)
	w := httptest.NewRecorder()
	v := map[string]string{"message": "Hello, world!"}
	err := Render(w, 200, v)
	assert.Nil(t, err)
	res := w.Result()
	defer res.Body.Close()
	assert.Equal(t, 200, res.StatusCode)
	var body map[string]string
	err = json.NewDecoder(res.Body).Decode(&body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello, world!", body["message"])
}

func TestReadJSON(t *testing.T) {
	log.SetOutput(io.Discard)
	data := map[string]interface{}{
		"name":         "Test1",
		"surName":      "Test2",
		"nation_id":    "12345678910",
		"phone_number": "+901234567890",
	}
	payload, err := json.Marshal(data)
	assert.Nil(t, err)
	req := httptest.NewRequest(http.MethodPost, "/readjson", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	// Test case 1: success
	user, err := ReadJSON(req)
	assert.Nil(t, err)
	assert.Equal(t, "Test1", user.Name)
	assert.Equal(t, "Test2", user.SurName)
	assert.Equal(t, "12345678910", user.Nation_Id)
	assert.Equal(t, "+901234567890", user.Phone_Number)
	// Test case 2: Empty request body
	req.Body = nil
	_, err = ReadJSON(req)
	assert.ErrorIs(t, err, ErrEmptyRequestBody)
	// Test case 3: Error while decoding JSON
	requestBody := `{"Invalid": "JSON"}`
	req.Body = io.NopCloser(strings.NewReader(requestBody))
	_, err = ReadJSON(req)
	assert.ErrorIs(t, err, ErrDecodeJSON)
	// Test case 4: Missing required params
	requestBody = `{"Name": "", "SurName": "test2", "Nation_Id": "", "Phone_Number": ""}`
	req.Body = io.NopCloser(strings.NewReader(requestBody))
	_, err = ReadJSON(req)
	assert.ErrorIs(t, err, ErrMissingRequiredParams)
}

func TestParseQuery_ReserveCar(t *testing.T) {
	log.SetOutput(io.Discard)
	req := &http.Request{
		URL: &url.URL{
			RawQuery: "id=123",
		},
	}
	// Test case 1: success
	id, err := ParseQuery_ReserveCar(req)
	assert.Nil(t, err)
	assert.Equal(t, 123, id)
	// Test case 2: Missing required parameters
	req = &http.Request{
		URL: &url.URL{},
	}
	_, err = ParseQuery_ReserveCar(req)
	assert.ErrorIs(t, err, ErrMissingRequiredParams)
	// Test case 3: Error while converting data
	req = &http.Request{
		URL: &url.URL{
			RawQuery: "id=test",
		},
	}
	_, err = ParseQuery_ReserveCar(req)
	assert.ErrorIs(t, err, ErrConvert)
}

func TestParseQuery_AvailableCars(t *testing.T) {
	log.SetOutput(io.Discard)
	req := &http.Request{
		URL: &url.URL{
			RawQuery: "receiver_date=2023-05-21&delivery_date=2023-05-23&location=1&time_start=10.00&time_end=18.00",
		},
	}
	// Test case 1: success
	actual, err := ParseQuery_AvailableCars(req)
	assert.Nil(t, err)
	expected := struct {
		DateInt   []int
		Location  int
		TimeStart string
		TimeEnd   string
	}{
		DateInt:   []int{0, 2},
		Location:  1,
		TimeStart: "10.00",
		TimeEnd:   "18.00",
	}
	assert.Equal(t, expected.DateInt, actual.DateInt)
	assert.Equal(t, expected.Location, actual.Location)
	assert.Equal(t, expected.TimeStart, actual.TimeStart)
	assert.Equal(t, expected.TimeEnd, actual.TimeEnd)

	// Test case 2: Missing required parameters
	req = &http.Request{
		URL: &url.URL{},
	}
	_, err = ParseQuery_AvailableCars(req)
	assert.ErrorIs(t, err, ErrMissingRequiredParams)

	// Test case 3: Error while converting data(date)
	req = &http.Request{
		URL: &url.URL{
			RawQuery: "receiver_date=2023-05-21&delivery_date=2023-05&location=1&time_start=10.00&time_end=18.00",
		},
	}
	_, err = ParseQuery_AvailableCars(req)
	assert.ErrorIs(t, err, ErrConvert)

	// Test case 3: Error while converting data(location)
	req = &http.Request{
		URL: &url.URL{
			RawQuery: "receiver_date=2023-05-21&delivery_date=2023-05-23&location=test&time_start=10.00&time_end=18.00",
		},
	}
	_, err = ParseQuery_AvailableCars(req)
	assert.ErrorIs(t, err, ErrConvert)
}
