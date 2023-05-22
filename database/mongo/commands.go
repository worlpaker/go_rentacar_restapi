package Mongo

import (
	"backend/models"
	Log "backend/pkg/helpers/log"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// ShowLocations retrieves a list of active locations from the database.
func (s *Server) ShowLocations() (result []models.Locations, err error) {
	collection := s.Client.Database("GODB").Collection("locations")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"active": true}
	cursor, err := collection.Find(ctx, filter)
	if Log.Err(err) {
		err = ErrFind
		return
	}
	if err = cursor.All(ctx, &result); Log.Err(err) {
		err = ErrCursorAll
		return
	} else if len(result) == 0 {
		err = ErrNotFoundResult
		return
	}
	return
}

// ShowReservedCars retrieves a list of reserved cars from the database.
func (s *Server) ShowReservedCars() (result []models.Cars, err error) {
	collection := s.Client.Database("GODB").Collection("cars")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"reserved": true}
	cursor, err := collection.Find(ctx, filter)
	if Log.Err(err) {
		err = ErrFind
		return
	}
	if err = cursor.All(ctx, &result); Log.Err(err) {
		err = ErrCursorAll
		return
	} else if len(result) == 0 {
		err = ErrNotFoundResult
		return
	}
	return
}

// ReserveCar reserves a car with the specified ID for the given user in the database.
func (s *Server) ReserveCar(id int, user *models.User) (err error) {
	collection := s.Client.Database("GODB").Collection("cars")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = UpdateUserBeforeMarshal(user)
	filter := bson.M{"id": id, "reserved": false}
	update := bson.M{"$set": bson.M{"reserved": true, "reserved_by": user}}
	result, err := collection.UpdateOne(ctx, filter, update)
	if Log.Err(err) {
		err = ErrUpdateOne
		return
	} else if result.ModifiedCount == 0 {
		err = ErrCarAlreadyReserved
		return
	}
	return
}

// AvailableCars retrieves a list of available cars based on the provided parameters.
func (s *Server) AvailableCars(data *models.AvailableCarsParams) (result []models.Cars, err error) {
	//1: find active offices by options
	offices_id, err := s.findActiveOffices(data)
	if Log.Err(err) {
		return
	}
	//2: list available unreserved cars
	collection := s.Client.Database("GODB").Collection("cars")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	filter := bson.M{"office_id": bson.M{"$in": offices_id}, "reserved": false}
	cursor, err := collection.Find(ctx, filter)
	if Log.Err(err) {
		err = ErrFind
		return
	}
	if err = cursor.All(ctx, &result); Log.Err(err) {
		err = ErrCursorAll
		return
	} else if len(result) == 0 {
		err = ErrNotFoundResult
		return
	}
	return
}

// findActiveOffices is a helper function for AvailableCars that finds active offices.
// Not exported.
func (s *Server) findActiveOffices(data *models.AvailableCarsParams) (id []int, err error) {
	var active_offices []models.Offices
	collection := s.Client.Database("GODB").Collection("offices")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	// if we want to search for multiple locations,
	// use: "location_id": bson.M{"$in": [array_of_loc_id]},
	filter := bson.M{
		"location_id":  data.Location,
		"working_days": bson.M{"$all": data.DateInt},
		"opening_hour": bson.M{"$lte": data.TimeStart},
		"closing_hour": bson.M{"$gte": data.TimeEnd},
	}
	cursor, err := collection.Find(ctx, filter)
	if Log.Err(err) {
		err = ErrFind
		return
	}
	if err = cursor.All(ctx, &active_offices); Log.Err(err) {
		err = ErrCursorAll
		return
	} else if len(active_offices) == 0 {
		err = ErrNoActiveOffices
		return
	}
	for i := range active_offices {
		id = append(id, active_offices[i].Id)
	}
	return
}
