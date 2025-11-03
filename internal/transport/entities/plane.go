package transport

type Plane struct {
	ID            string  // уникальный идентификатор (например, UUID)
	Model         string  // марка и модель
	FuelLvl       int     // уровень топлива (0–100)
	DirtLvl       int     // загрязнение (0–100)
	Mileage       int     // пробег в км
	Condition     int     // состояние (0–100, где 100 — идеал)
	Rented        bool    // сейчас в аренде?
	PricePerDay   float64 // цена за день аренды
	LastServiceKm int     // пробег последнего ТО
}

// type Helicopter struct {

// }

// Функция конструктор для создания самолета
func NewPlane() (*Plane, error) {

	return nil, nil
}
