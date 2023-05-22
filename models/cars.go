package models

type Cars struct {
	Base
	Id           int    `json:"id"`
	Vendor       string `json:"vendor"`
	Fuel         string `json:"fuel"`
	Transmission string `json:"transmission"`
	Name         string `json:"name"`
	Office_Id    int    `json:"office_id"`
	Reserved     bool   `json:"reserved"`
	Reserved_By  User   `json:"reserved_by"`
}
