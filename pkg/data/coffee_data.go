package data

import "coffee-goes-fullstack/pkg/models"

var espresso = models.Coffee{
	ID:    "1",
	Water: 250,
	Milk:  0,
	Beans: 16,
	Cost:  4,
}

var latte = models.Coffee{
	ID:    "2",
	Water: 350,
	Milk:  75,
	Beans: 20,
	Cost:  7,
}

var cappuccino = models.Coffee{
	ID:    "3",
	Water: 200,
	Milk:  100,
	Beans: 12,
	Cost:  6,
}

var Coffees = []models.Coffee{espresso, latte, cappuccino}
