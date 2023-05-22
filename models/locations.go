package models

type Locations struct {
	Base
	Id     int    `json:"Id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}
