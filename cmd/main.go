package main

import (
	"fmt"
	"log"

	"Transport_rental/internal/common"
	"Transport_rental/internal/genid"
	"Transport_rental/internal/tariffs"
	transport "Transport_rental/internal/transport/entities"
)

func main() {
	genid.NewID("user", genid.UuidGen{})   // create user ID
	genid.NewID("car", genid.ShortIDGen{}) // create transport ID

	genid.ShowAllID()    // show all IDs
	common.ShowLogsErr() // show all logs with errors

	// допустим, всего в автопарке 10 машин, сейчас арендованы 3
	totalCars := 10
	rentedCars := 3

	// создаём новую машину
	car, err := transport.NewCar("Toyota Camry", "Comfort", tariffs.ComfortCar, rentedCars, totalCars)
	if err != nil {
		log.Fatalf("Ошибка при создании машины: %v", err)
	}

	// выводим информацию о машине
	fmt.Printf("Новая машина создана:\n")
	fmt.Printf("ID: %s\n", car.ID)
	fmt.Printf("Модель: %s\n", car.Model)
	fmt.Printf("Класс: %s\n", car.Class)
	fmt.Printf("Тип: %s\n", car.VehicleType)
	fmt.Printf("Цена за час: %.2f ₽\n", car.PricePerHour)
	fmt.Printf("Цена за день: %.2f ₽\n", car.PricePerDay)
	fmt.Printf("Топливо: %d%%, Состояние: %d%%\n", car.FuelLvl, car.Condition)
}
