package interfaces

// Transport — базовый интерфейс для всех видов транспорта.
type Transport interface {
	LvlVehicleWear() // износ
	Repair()         // ремонт
}

type CheckCondition interface {
	// TODO:
}