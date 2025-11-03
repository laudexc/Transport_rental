package genid

import "math/rand/v2"

func NewParams() (fuellvl, dirtlvl, mileage, cond, lastServ int, rented bool) {
	fuellvl = rand.IntN(101)
	dirtlvl = rand.IntN(101)
	mileage = rand.IntN(200001)
	cond = rand.IntN(101)
	lastServ = rand.IntN(50001)

	randRen := rand.IntN(2)
	switch randRen {
	case 0:
		rented = false
	default:
		rented = true
	}

	return fuellvl, dirtlvl, mileage, cond, lastServ, rented
}
