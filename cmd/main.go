package main

import (
	"Transport_rental/internal/genid"
	"Transport_rental/internal/transport/entities"
)

func main() {
	transportID := NewID("car", ShortIDGen{})
	userID := NewID("user", UuidGen{})
}
