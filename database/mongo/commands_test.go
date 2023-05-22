package Mongo

import (
	"backend/models"
	"io"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func fakeMongoAuth(m *mongo.Client) *Server {
	return &Server{Client: m}
}

func TestShowLocations(t *testing.T) {
	log.SetOutput(io.Discard)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		s := fakeMongoAuth(mt.Client)
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
		actual_result, err := s.ShowLocations()
		assert.Nil(t, err)
		assert.Equal(t, test_data, actual_result)
	})
	mt.Run("no_data", func(mt *mtest.T) {
		s := fakeMongoAuth(mt.Client)
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "foo.bar", mtest.FirstBatch))
		actualResult, err := s.ShowLocations()
		assert.ErrorIs(t, err, ErrNotFoundResult)
		assert.Empty(t, actualResult)
	})

	mt.Run("find_error", func(mt *mtest.T) {
		s := fakeMongoAuth(mt.Client)
		mt.AddMockResponses(
			mtest.CreateCommandErrorResponse(mtest.CommandError{Message: "find error"}),
		)
		actualResult, err := s.ShowLocations()
		assert.ErrorIs(t, err, ErrFind)
		assert.Empty(t, actualResult)
	})

	mt.Run("cursor_all_error", func(mt *mtest.T) {
		s := fakeMongoAuth(mt.Client)
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "id", Value: 1},
			{Key: "name", Value: "Istanbul"},
			{Key: "Active", Value: true},
		}), mtest.CreateCommandErrorResponse(mtest.CommandError{Message: "cursor_all_error"}))
		actualResult, err := s.ShowLocations()
		assert.ErrorIs(t, err, ErrCursorAll)
		assert.Empty(t, actualResult)
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
			Vendor:       "Somevendor",
			Fuel:         "Petrol",
			Transmission: "Manual",
			Name:         "Ford Focus",
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
			Vendor:       "Somevendor",
			Fuel:         "Petrol",
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

	mt.Run("success", func(mt *mtest.T) {
		s := fakeMongoAuth(mt.Client)
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
		actual_result, err := s.ShowReservedCars()
		assert.Nil(t, err)
		assert.Equal(t, test_data, actual_result)
	})
	mt.Run("no_data", func(mt *mtest.T) {
		s := fakeMongoAuth(mt.Client)
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "foo.bar", mtest.FirstBatch))
		actualResult, err := s.ShowReservedCars()
		assert.ErrorIs(t, err, ErrNotFoundResult)
		assert.Empty(t, actualResult)
	})

	mt.Run("find_error", func(mt *mtest.T) {
		s := fakeMongoAuth(mt.Client)
		mt.AddMockResponses(
			mtest.CreateCommandErrorResponse(mtest.CommandError{Message: "find error"}),
		)
		actualResult, err := s.ShowReservedCars()
		assert.ErrorIs(t, err, ErrFind)
		assert.Empty(t, actualResult)
	})

	mt.Run("cursor_all_error", func(mt *mtest.T) {
		s := fakeMongoAuth(mt.Client)
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
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
		}), mtest.CreateCommandErrorResponse(mtest.CommandError{Message: "cursor_all_error"}))
		actualResult, err := s.ShowReservedCars()
		assert.Empty(t, actualResult)
		assert.ErrorIs(t, err, ErrCursorAll)
	})
}

func TestReserveCar(t *testing.T) {
	log.SetOutput(io.Discard)
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	test_cardata := &models.Cars{
		Reserved:     false,
		Id:           1,
		Vendor:       "Somevendor",
		Fuel:         "Petrol",
		Transmission: "Automatic",
		Name:         "Porsche Panamera",
		Office_Id:    1,
	}
	test_userdata := &models.User{
		Name:         "test2",
		SurName:      "test2",
		Nation_Id:    "12345678910",
		Phone_Number: "+901234567890",
	}
	mt.Run("success", func(mt *mtest.T) {
		s := fakeMongoAuth(mt.Client)
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "nModified", Value: 1},
		})
		err := s.ReserveCar(test_cardata.Id, test_userdata)
		assert.Nil(t, err)
	})
	mt.Run("update_error", func(mt *mtest.T) {
		s := fakeMongoAuth(mt.Client)
		mt.AddMockResponses(
			mtest.CreateCommandErrorResponse(mtest.CommandError{Message: "find error"}),
		)
		err := s.ReserveCar(1, test_userdata)
		assert.ErrorIs(t, err, ErrUpdateOne)
	})

	mt.Run("car_reserved_error", func(mt *mtest.T) {
		s := fakeMongoAuth(mt.Client)
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "foo.bar", mtest.FirstBatch))
		err := s.ReserveCar(1, test_userdata)
		assert.ErrorIs(t, err, ErrCarAlreadyReserved)
	})

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
			Vendor:       "Testvendor",
			Fuel:         "Petrol",
			Transmission: "Manual",
			Name:         "Hyundai i30",
			Office_Id:    1,
		},
	}
	test_params := &models.AvailableCarsParams{
		DateInt:   []int{1, 2},
		Location:  1,
		TimeStart: "10.00",
		TimeEnd:   "18.00",
	}

	mt.Run("success", func(mt *mtest.T) {
		s := fakeMongoAuth(mt.Client)
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

		actual_result, err := s.AvailableCars(test_params)
		assert.Nil(t, err)
		assert.Equal(t, test_cardata, actual_result)
	})
	mt.Run("offices_not_active", func(mt *mtest.T) {
		s := fakeMongoAuth(mt.Client)
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "foo.bar", mtest.FirstBatch))
		actualResult, err := s.AvailableCars(test_params)
		assert.ErrorIs(t, err, ErrNoActiveOffices)
		assert.Empty(t, actualResult)
	})

	mt.Run("offices_find_error", func(mt *mtest.T) {
		s := fakeMongoAuth(mt.Client)
		mt.AddMockResponses(
			mtest.CreateCommandErrorResponse(mtest.CommandError{Message: "offices find error"}),
		)
		actualResult, err := s.AvailableCars(test_params)
		assert.ErrorIs(t, err, ErrFind)
		assert.Empty(t, actualResult)
	})

	mt.Run("offices_cursor_all_error", func(mt *mtest.T) {
		s := fakeMongoAuth(mt.Client)
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "reserved", Value: test_cardata[0].Reserved},
			{Key: "id", Value: test_cardata[0].Id},
			{Key: "vendor", Value: test_cardata[0].Vendor},
			{Key: "fuel", Value: test_cardata[0].Fuel},
			{Key: "transmission", Value: test_cardata[0].Transmission},
			{Key: "name", Value: test_cardata[0].Name},
			{Key: "office_id", Value: test_cardata[0].Office_Id},
		}), mtest.CreateCommandErrorResponse(mtest.CommandError{Message: "offices_cursor_all_error"}))
		actualResult, err := s.AvailableCars(test_params)
		assert.ErrorIs(t, err, ErrCursorAll)
		assert.Empty(t, actualResult)
	})

	mt.Run("availablecars_find_error", func(mt *mtest.T) {
		s := fakeMongoAuth(mt.Client)
		test1 := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "id", Value: test_officedata.Id},
			{Key: "location_id", Value: test_officedata.Location_Id},
			{Key: "opening_hour", Value: test_officedata.Opening_Hour},
			{Key: "closing_hour", Value: test_officedata.Closing_Hour},
			{Key: "working_days", Value: test_officedata.Working_Days},
		})
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(test1, killCursors)
		mt.AddMockResponses(
			mtest.CreateCommandErrorResponse(mtest.CommandError{Message: "availablecars_find_error"}),
		)
		actualResult, err := s.AvailableCars(test_params)
		assert.ErrorIs(t, err, ErrFind)
		assert.Empty(t, actualResult)
	})

	mt.Run("availablecars_cursor_error", func(mt *mtest.T) {
		s := fakeMongoAuth(mt.Client)
		test1 := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "id", Value: test_officedata.Id},
			{Key: "location_id", Value: test_officedata.Location_Id},
			{Key: "opening_hour", Value: test_officedata.Opening_Hour},
			{Key: "closing_hour", Value: test_officedata.Closing_Hour},
			{Key: "working_days", Value: test_officedata.Working_Days},
		})
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(test1, killCursors)
		mt.AddMockResponses(test1,
			mtest.CreateCommandErrorResponse(mtest.CommandError{Message: "availablecars_cursor_error"}),
		)
		actualResult, err := s.AvailableCars(test_params)
		assert.ErrorIs(t, err, ErrCursorAll)
		assert.Empty(t, actualResult)
	})

	mt.Run("availablecars_no_data", func(mt *mtest.T) {
		s := fakeMongoAuth(mt.Client)
		test1 := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "id", Value: test_officedata.Id},
			{Key: "location_id", Value: test_officedata.Location_Id},
			{Key: "opening_hour", Value: test_officedata.Opening_Hour},
			{Key: "closing_hour", Value: test_officedata.Closing_Hour},
			{Key: "working_days", Value: test_officedata.Working_Days},
		})
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(test1, killCursors)
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "foo.bar", mtest.FirstBatch))
		actualResult, err := s.AvailableCars(test_params)
		assert.ErrorIs(t, err, ErrNotFoundResult)
		assert.Empty(t, actualResult)
	})
}
