package transport

import (
	"fmt"
	"strings"

	"Transport_rental/internal/common"
	"Transport_rental/internal/genid"
	"Transport_rental/internal/tariffs"
)

type Car struct {
	ID            string
	Model         string
	Class         string
	VehicleType   tariffs.VehicleType
	FuelLvl       int
	DirtLvl       int
	Mileage       int
	Condition     int
	LastServiceKm int
	Rented        bool
	PricePerHour  float64
	PricePerDay   float64
}

// Constructor function for creating a car
func NewCar(model, class string, vType tariffs.VehicleType, rentedCount, totalCount int) (*Car, error) {
	if strings.TrimSpace(model) == "" || strings.TrimSpace(class) == "" {
		return nil, fmt.Errorf("invalid car parameters: %s", common.EmptyIdOrModel)
	}

	id := genid.NewID("Car", genid.ShortIDGen{})
	fuellvl, dirtlvl, mileage, cond, lastServ, rented := genid.NewParams()

	priceH, priceD := tariffs.CalculatePrice(vType, rentedCount, totalCount)
	return &Car{
		ID:            id,
		Model:         model,
		Class:         class,
		VehicleType:   vType,
		FuelLvl:       fuellvl,
		DirtLvl:       dirtlvl,
		Mileage:       mileage,
		Condition:     cond,
		LastServiceKm: lastServ,
		Rented:        rented,
		PricePerHour:  priceH,
		PricePerDay:   priceD,
	}, nil
}

// Методы состояния

// Минимум:

// Rent() — арендовать транспорт (меняет флаг IsRented = true);

// Return() — вернуть;

// IsAvailable() — возвращает bool, свободен ли транспорт;

// Info() или String() — для удобного вывода информации.

// type ElectricCar struct {
// 	Car                 // встраиваем обычный Car
// 	BatteryCapacity int // ёмкость батареи (в кВт·ч)
// 	ChargeLevel     int // уровень заряда (0–100)
// 	ChargingTime    int // время до полного заряда (в минутах)
// }

// func NewECar() (*ElecticCar, error) {

// }
