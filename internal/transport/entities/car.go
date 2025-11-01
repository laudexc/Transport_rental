package transport

// Car — структура автомобиля.
type Car struct {
	ID            string  // уникальный идентификатор (например, UUID)
	Model         string  // марка и модель
	FuelLvl       int     // уровень топлива (0–100)
	DirtLvl       int     // загрязнение (0–100)
	TotalDistance int     // пробег в км
	Condition     int     // состояние (0–100, где 100 — идеал)
	Rented        bool    // сейчас в аренде?
	PricePerDay   float64 // цена за день аренды
	LastServiceKm int     // пробег последнего ТО
}

// // ElectricCar — структура электромобиля.
// type ElectricCar struct {
// 	Car                 // встраиваем обычный Car
// 	BatteryCapacity int // ёмкость батареи (в кВт·ч)
// 	ChargeLevel     int // уровень заряда (0–100)
// 	ChargingTime    int // время до полного заряда (в минутах)
// }

// Функция конструктор для создания машины
func NewCar() (*Car, error) {

	return nil, nil
}
