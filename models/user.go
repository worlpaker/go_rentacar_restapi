package models

type User struct {
	Base
	Name         string `json:"name"`
	SurName      string `json:"surname"`
	Nation_Id    string `json:"nation_id"`
	Phone_Number string `json:"phone_number"`
}
