package models

type Offices struct {
	Base
	Id           int    `json:"id"`
	Vendor       string `json:"Vendor"`
	Location_Id  int    `json:"location_id"`
	Opening_Hour string `json:"opening_hour"`
	Closing_Hour string `json:"closing_hour"`
	Working_Days []int  `json:"working_days"`
}
