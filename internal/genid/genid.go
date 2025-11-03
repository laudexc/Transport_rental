package genid

import (
	"Transport_rental/internal/common"
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/google/uuid"
)

type Generator interface {
	generate() (string, error)
}

var ID = make(map[string]string)

type UuidGen struct{}

func (u UuidGen) generate() (string, error) {
	for {
		id, err := uuid.NewUUID()
		if _, exist := ID[id.String()]; !exist { // collision check
			return id.String(), err
		}
	}
}

type ShortIDGen struct{}

func (sig ShortIDGen) generate() (string, error) {
	for {
		num := rand.IntN(10000)
		ShID := fmt.Sprintf("Transport-%d", num)
		if _, exist := ID[ShID]; !exist { // collision check
			return ShID, nil
		}
	}

}

// функция для работы со всеми генераторами
func NewID(entityType string, gen Generator) string {
	id, err := gen.generate()
	if err != nil {
		common.LogError("genid", err)
		return ""
	}

	ID[entityType] = id
	log.Printf("[genid] ID for %s successfully created: %s\n", entityType, id)
	return id
}

func ShowAllID() {
	fmt.Printf("***---***\nAll users IDs:\n")
	num := 1
	for entityType, id := range ID {
		fmt.Printf("%d. Entity Type: %s, ID: %s\n", num, entityType, id)
		num += 1
	}
	fmt.Println("***---***")

	// To check error logs, see common/errors ShowLogsErr()
}
