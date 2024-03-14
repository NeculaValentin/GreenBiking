package main

import "time"

type StazioneServizio struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	NomeImpianto  string  `json:"nomeImpianto"`
	Address       string  `json:"address"`
	Brand         string  `json:"brand"`
	Fuels         []Fuels `json:"fuels"`
	PhoneNumber   string  `json:"phoneNumber"`
	Email         string  `json:"email"`
	Website       string  `json:"website"`
	Company       string  `json:"company"`
	Services      []any   `json:"services"`
	Orariapertura []any   `json:"orariapertura"`
}

type Fuels struct {
	ID            int       `json:"id"`
	Price         float64   `json:"price"`
	Name          string    `json:"name"`
	FuelID        int       `json:"fuelId"`
	IsSelf        bool      `json:"isSelf"`
	ServiceAreaID int       `json:"serviceAreaId"`
	InsertDate    time.Time `json:"insertDate"`
	ValidityDate  time.Time `json:"validityDate"`
}

const (
	ALL = iota
	PETROL
	DIESEL
	METHANE
)
