package tariffs

import "time"

type VehicleType string

const (
	EconomyCar  VehicleType = "Economy car"
	ComfortCar  VehicleType = "Comfort car"
	BusinessCar VehicleType = "Business car"
	PremiumCar  VehicleType = "Premium car"

	EconomyBike VehicleType = "Economy bike"
	ComfortBike VehicleType = "Comfort bike"

	EconomyBoat VehicleType = "Economy boat"
	ComfortBoat VehicleType = "Comfort boat"
	PremiumBoat VehicleType = "Premium boat"

	DefaultPlane VehicleType = "Default plane"
	BusinessJet  VehicleType = "Business jet"
)

var BasePrices = map[VehicleType]struct {
	PerHour float64
	PerDay  float64
}{
	EconomyCar:  {PerHour: 350, PerDay: 2200},
	ComfortCar:  {PerHour: 550, PerDay: 3500},
	BusinessCar: {PerHour: 900, PerDay: 6000},
	PremiumCar:  {PerHour: 1400, PerDay: 9500},

	EconomyBike: {PerHour: 100, PerDay: 600},
	ComfortBike: {PerHour: 150, PerDay: 950},

	EconomyBoat: {PerHour: 1500, PerDay: 10000},
	ComfortBoat: {PerHour: 2500, PerDay: 16000},
	PremiumBoat: {PerHour: 4000, PerDay: 25000},

	DefaultPlane: {PerHour: 15000, PerDay: 100000},
	BusinessJet:  {PerHour: 35000, PerDay: 250000},
}

// A function that takes into account discounts, seasonality, number of cars, etc.
func CalculatePrice(vType VehicleType, rentedCount, totalCount int) (hourly, daily float64) {
	base := BasePrices[vType]

	// dynamic pricing depending on demand
	overallLoad := float64(rentedCount) / float64(totalCount)
	if overallLoad >= 0.8 { // if more than 80% of the cars are rented
		base.PerHour *= 1.25
		base.PerDay *= 1.25
	}

	// holiday discount
	if isHoliday(time.Now()) {
		base.PerHour *= 0.9
		base.PerDay *= 0.9
	}

	return base.PerHour, base.PerDay
}

func isHoliday(t time.Time) bool {
	month := t.Month()
	day := t.Day()

	// New Year and New Year holidays (1-6 & 8 january)
	if month == time.January && (day >= 1 && day <= 6 || day == 8) {
		return true
	}
	// Christmas (January 7)
	if month == time.January && day == 7 {
		return true
	}
	// Defender of the Fatherland Day (February 23)
	if month == time.February && day == 23 {
		return true
	}
	// International Women's Day (8 March)
	if month == time.March && day == 8 {
		return true
	}
	// Spring and Labor Day (May 1)
	if month == time.May && day == 1 {
		return true
	}
	// Victory Day (May 9)
	if month == time.May && day == 9 {
		return true
	}
	// Russia Day (June 12)
	if month == time.June && day == 12 {
		return true
	}
	// National Unity Day (November 4)
	if month == time.November && day == 4 {
		return true
	}

	return false
}
