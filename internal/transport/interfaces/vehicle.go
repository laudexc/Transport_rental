package interfaces

// Transport — базовый интерфейс для всех видов транспорта.
type Transport interface {
	Drive()           // движение
	Stop()            // остановка
	WearVehicleLvl()  // износ
	Repair()          // ремонт
}
