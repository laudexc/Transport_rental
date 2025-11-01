package main

import (
	"fmt"
)

// TODO:func (b *Car) WearVehicleLvl() {}

// TODO: func (b *Car) Repair() {}

func (c *Car) Drive(distance int) {
	FuelFlow := distance * FuelPerKm

	// Если нет топлива
	if c.FuelLvl == 0 {
		fmt.Println("Топливо на нуле, машина не может завестись")
		return
	}

	// Если топлива не хватает на всю поездку
	if FuelFlow > c.FuelLvl {
		maxDistance := c.FuelLvl / FuelFlow
		fmt.Printf("Недостаточно топлива для %d км. Машина сможет проехать только %d км и заглохнет!\n", distance, maxDistance)
		c.FuelLvl = 0 // После поездки топливо на нуле
		return
	}

	// Если топлива хватит, совершаем поездку и снимаем требуемое количество топлива
	c.FuelLvl -= FuelFlow
	fmt.Printf("Топлива хватает, машина проехала %d км. Остаток топлива: %d%%\n", distance, c.FuelLvl)
}

func (c *Car) Checkfuel() {
	fmt.Printf("Уровень топлива - %v%%\n", c.FuelLvl)
}

// TODO: доделать func (c *Car) Refuel(amount int) {}

func (c Car) Stop() {
	fmt.Println("Машина остановилась")
}

// Все для Bike
type ElecticBike struct {
	FuelType   string
	BatteryLvl int
}

const BatteryPerKm = 1

func (eb *ElecticBike) Drive(distance int) {
	BatteryFlow := distance * BatteryPerKm

	if eb.BatteryLvl == 0 {
		fmt.Println("Аккумулятор байка разряжен или отсутствует , байк не заведётся!")
		return
	}

	if BatteryFlow > eb.BatteryLvl {
		MaxDistance := eb.BatteryLvl / BatteryFlow
		fmt.Printf("Недостаточно заряда для %d км. Байк проедет %d км и отключится\n", distance, MaxDistance)
		eb.BatteryLvl = 0 // Байк разрядился
		return
	}

	if BatteryFlow < eb.BatteryLvl {
		eb.BatteryLvl -= BatteryFlow
		fmt.Printf("Заряда хватило на поездку, байк проехал %d км. Остаток заряда: %d%%\n", distance, eb.BatteryLvl)
	}
}

func (eb *ElecticBike) CheckBattery() {
	fmt.Printf("Уровень заряда байка: %v%%\n", eb.BatteryLvl)
}

// TODO: func (eb *ElecticBike) RechargeBattery(amount int) {}

// TODO: func (eb *ElecticBike) WearVehicleLvl() {}

// TODO:func (eb *ElecticBike) Repair() {}

func (b *ElecticBike) Stop() {
	fmt.Println("Байк остановился")
}

// Все для Boat
type Boat struct {
	EngineOn bool
	FuelType string
}

// TODO: func (b *Boat) WearVehicleLvl() {}

// TODO: func (b *Boat) Repair() {}

func (bt *Boat) StartEngine() bool {
	fmt.Println("Двигатель лодки запущен")
	bt.EngineOn = true
	return true
}

func (bt *Boat) Drive() { // TODO: доделать
	if bt.EngineOn == true {
		fmt.Println("Мотор заведен, лодка начинает движение")
	} else {
		fmt.Println("Мотор не заведен, лодка не может начать движение")
	}
}

func (bt *Boat) Stop() {
	fmt.Println("Лодка остановилась")
}

func StartAllTransport(vehicles []Transport) {

	for _, v := range vehicles {
		v.Drive()
		v.Stop()
	}
}

// TODO: func TrafficSimulator() {}

func main() {
	car := Car{FuelLevel: 90}
	//electicbike := ElecticBike{HasRider: true}
	//boat := Boat{EngineOn: false}

	fmt.Println("Введите количество км, которое планируете проехать:")
	var distance int
	_, err := fmt.Scanf("%d", &distance)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	car.Drive(distance)

	/* TODO: ?? vehicles := []Transport{&car, &electicbike, &boat}
	StartAllTransport(vehicles)
	*/
}
