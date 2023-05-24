package routers

import (
	"backend/api/handlers"
	"backend/database"
	Mongo "backend/database/mongo"
	"backend/models"
	"bytes"
	"io"
	"log"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

// Handlers API Tests

func executeRequest(r *http.Request, s *http.ServeMux) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)
	return w
}

func fakeMongoAuth(m *mongo.Client) *Mongo.Server {
	return &Mongo.Server{Client: m}
}

func fakeNewServer(m *mongo.Client) *Server {
	log.SetOutput(io.Discard)
	r := http.NewServeMux()
	d := fakeMongoAuth(m)
	s := &Server{
		Router: r,
		Handlers: &handlers.Server{
			DB: &database.DB{
				Mongo: d,
			},
		},
	}
	s.SetupRouters()
	return s
}

func convertJSONToBuf(t *testing.T, data interface{}) *bytes.Buffer {
	dataBuf := new(bytes.Buffer)
	err := json.NewEncoder(dataBuf).Encode(data)
	assert.Nil(t, err)
	return dataBuf
}

func TestShowLocations(t *testing.T) {
	log.SetOutput(io.Discard)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	test_data := []models.Locations{
		{
			Id:     1,
			Name:   "Istanbul",
			Active: true,
		},
		{
			Id:     3,
			Name:   "Paris",
			Active: true,
		},
	}
	mt.Run("success_and_error", func(mt *mtest.T) {
		s := fakeNewServer(mt.Client)
		test1 := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "id", Value: test_data[0].Id},
			{Key: "name", Value: test_data[0].Name},
			{Key: "Active", Value: test_data[0].Active},
		})
		test2 := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, bson.D{
			{Key: "id", Value: test_data[1].Id},
			{Key: "name", Value: test_data[1].Name},
			{Key: "Active", Value: test_data[1].Active},
		})
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(test1, test2, killCursors)
		req, err := http.NewRequest("GET", "/api/locations/show", nil)
		assert.Nil(t, err)
		resp := executeRequest(req, s.Router)
		var actual_result []models.Locations
		err = json.NewDecoder(resp.Body).Decode(&actual_result)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, test_data, actual_result)
		// test for the error
		req, err = http.NewRequest("GET", "/api/locations/show", nil)
		assert.Nil(t, err)
		resp = executeRequest(req, s.Router)
		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		assert.Equal(t, "failed to find data\n", resp.Body.String())
	})

}

func TestShowReservedCars(t *testing.T) {
	log.SetOutput(io.Discard)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	test_data := []models.Cars{
		{
			Reserved:     true,
			Id:           1,
			Vendor:       "test",
			Fuel:         "Petrol",
			Transmission: "Manual",
			Name:         "Fiat Egea",
			Office_Id:    1,
			Reserved_By: models.User{
				Name:         "test1",
				SurName:      "test1",
				Nation_Id:    "12345678910",
				Phone_Number: "+901234567890",
			},
		},
		{
			Reserved:     true,
			Id:           2,
			Vendor:       "test2",
			Fuel:         "Diesel",
			Transmission: "Automatic",
			Name:         "Volvo V90",
			Office_Id:    4,
			Reserved_By: models.User{
				Name:         "test2",
				SurName:      "test2",
				Nation_Id:    "12345678910",
				Phone_Number: "+901234567890",
			},
		},
	}
	mt.Run("success_and_error", func(mt *mtest.T) {
		s := fakeNewServer(mt.Client)
		test1 := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "reserved", Value: test_data[0].Reserved},
			{Key: "id", Value: test_data[0].Id},
			{Key: "vendor", Value: test_data[0].Vendor},
			{Key: "fuel", Value: test_data[0].Fuel},
			{Key: "transmission", Value: test_data[0].Transmission},
			{Key: "name", Value: test_data[0].Name},
			{Key: "office_id", Value: test_data[0].Office_Id},
			{Key: "reserved_by", Value: bson.D{
				{Key: "name", Value: test_data[0].Reserved_By.Name},
				{Key: "surname", Value: test_data[0].Reserved_By.SurName},
				{Key: "nation_id", Value: test_data[0].Reserved_By.Nation_Id},
				{Key: "phone_number", Value: test_data[0].Reserved_By.Phone_Number},
			}},
		})
		test2 := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, bson.D{
			{Key: "reserved", Value: test_data[1].Reserved},
			{Key: "id", Value: test_data[1].Id},
			{Key: "vendor", Value: test_data[1].Vendor},
			{Key: "fuel", Value: test_data[1].Fuel},
			{Key: "transmission", Value: test_data[1].Transmission},
			{Key: "name", Value: test_data[1].Name},
			{Key: "office_id", Value: test_data[1].Office_Id},
			{Key: "reserved_by", Value: bson.D{
				{Key: "name", Value: test_data[1].Reserved_By.Name},
				{Key: "surname", Value: test_data[1].Reserved_By.SurName},
				{Key: "nation_id", Value: test_data[1].Reserved_By.Nation_Id},
				{Key: "phone_number", Value: test_data[1].Reserved_By.Phone_Number},
			}},
		})
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(test1, test2, killCursors)
		req, err := http.NewRequest("GET", "/api/cars/showreservedcars", nil)
		assert.Nil(t, err)
		resp := executeRequest(req, s.Router)
		var actual_result []models.Cars
		err = json.NewDecoder(resp.Body).Decode(&actual_result)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, test_data, actual_result)
		//simple test for the error
		req, err = http.NewRequest("GET", "/api/cars/showreservedcars", nil)
		assert.Nil(t, err)
		resp = executeRequest(req, s.Router)
		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		assert.Equal(t, "failed to find data\n", resp.Body.String())
	})

}

func TestReserveCar(t *testing.T) {
	log.SetOutput(io.Discard)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	test_userdata := &models.User{
		Name:         "test1",
		SurName:      "test2",
		Nation_Id:    "12345678910",
		Phone_Number: "+901234567890",
	}
	testData := []struct {
		name          string
		db_active     bool
		mockResponses bson.D
		requestBody   io.Reader
		requestURL    string
		expectedCode  int
		expectedBody  string
	}{
		{
			name: "success",
			mockResponses: bson.D{
				{Key: "ok", Value: 1},
				{Key: "nModified", Value: 1},
			},
			requestBody:  convertJSONToBuf(t, test_userdata),
			requestURL:   "/api/cars/reserve?id=1",
			expectedCode: http.StatusCreated,
			expectedBody: "\"successfully reserved car 1\"\n",
		},
		{
			name:          "error",
			mockResponses: nil,
			requestBody:   convertJSONToBuf(t, test_userdata),
			requestURL:    "/api/cars/reserve",
			expectedCode:  http.StatusBadRequest,
			expectedBody:  "missing required parameters\n",
		},
		{
			name:          "error",
			mockResponses: nil,
			requestBody:   nil,
			requestURL:    "/api/cars/reserve?id=1",
			expectedCode:  http.StatusBadRequest,
			expectedBody:  "empty request body\n",
		},
		{
			name:          "error",
			mockResponses: mtest.CreateWriteErrorsResponse(),
			requestBody:   convertJSONToBuf(t, test_userdata),
			requestURL:    "/api/cars/reserve?id=1",
			expectedCode:  http.StatusInternalServerError,
			expectedBody:  "the car is already reserved\n",
		},
		{
			name:          "error",
			mockResponses: nil,
			requestBody:   convertJSONToBuf(t, test_userdata),
			requestURL:    "/api/cars/reserve?id=1",
			expectedCode:  http.StatusInternalServerError,
			expectedBody:  "failed to update data\n",
		},
	}
	for _, k := range testData {
		mt.Run(k.name, func(mt *mtest.T) {
			s := fakeNewServer(mt.Client)
			req, err := http.NewRequest("POST", k.requestURL, k.requestBody)
			assert.Nil(t, err)
			mt.AddMockResponses(k.mockResponses)
			resp := executeRequest(req, s.Router)
			assert.Equal(t, k.expectedCode, resp.Code)
			assert.Equal(t, k.expectedBody, resp.Body.String())
		})
	}
}

func TestAvailableCars(t *testing.T) {
	log.SetOutput(io.Discard)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	test_officedata := models.Offices{
		Id:           1,
		Location_Id:  1,
		Opening_Hour: "08.00",
		Closing_Hour: "18.00",
		Working_Days: []int{1, 2, 3, 4, 5},
	}
	test_cardata := []models.Cars{
		{
			Reserved:     false,
			Id:           1,
			Vendor:       "test",
			Fuel:         "Diesel",
			Transmission: "Hybrid",
			Name:         "Peugeot 408",
			Office_Id:    1,
		},
	}
	mt.Run("success", func(mt *mtest.T) {
		s := fakeNewServer(mt.Client)

		test1 := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "id", Value: test_officedata.Id},
			{Key: "location_id", Value: test_officedata.Location_Id},
			{Key: "opening_hour", Value: test_officedata.Opening_Hour},
			{Key: "closing_hour", Value: test_officedata.Closing_Hour},
			{Key: "working_days", Value: test_officedata.Working_Days},
		})
		test2 := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "reserved", Value: test_cardata[0].Reserved},
			{Key: "id", Value: test_cardata[0].Id},
			{Key: "vendor", Value: test_cardata[0].Vendor},
			{Key: "fuel", Value: test_cardata[0].Fuel},
			{Key: "transmission", Value: test_cardata[0].Transmission},
			{Key: "name", Value: test_cardata[0].Name},
			{Key: "office_id", Value: test_cardata[0].Office_Id},
		})
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(test1, killCursors)
		mt.AddMockResponses(test2, killCursors)
		url := "/api/cars/available?time_start=09.00&time_end=21.00&receiver_date=2023-05-03&delivery_date=2023-05-08&location=1"
		req, err := http.NewRequest("GET", url, nil)
		assert.Nil(t, err)
		resp := executeRequest(req, s.Router)
		var actual_result []models.Cars
		err = json.NewDecoder(resp.Body).Decode(&actual_result)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, test_cardata, actual_result)
		// test for the error 1: parameter
		url = "/api/cars/available?time_start=09.00&receiver_date=2023-05-03&delivery_date=2023-05-08"
		req, err = http.NewRequest("GET", url, nil)
		assert.Nil(t, err)
		resp = executeRequest(req, s.Router)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Equal(t, "missing required parameters\n", resp.Body.String())
		//test for the error 2: db error
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse())
		url = "/api/cars/available?time_start=09.00&time_end=09.00&receiver_date=2023-05-03&delivery_date=2023-05-08&location=1"
		req, err = http.NewRequest("GET", url, nil)
		assert.Nil(t, err)
		resp = executeRequest(req, s.Router)
		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		assert.Equal(t, "failed to find data\n", resp.Body.String())
	})
}

// Test for Routers

func TestUseMw(t *testing.T) {
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	req := httptest.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	handler := UseMw(mockHandler)
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "Hello, World!", resp.Body.String())
}
