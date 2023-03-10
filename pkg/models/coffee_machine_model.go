package models

type CoffeeMachine struct {
	Water int `json:"water"`
	Milk  int `json:"milk"`
	Beans int `json:"beans"`
	Cups  int `json:"cups"`
	Money int `json:"cost"`
}
