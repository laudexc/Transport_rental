package genid

import (
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/google/uuid"
)

type Generator interface {
	generate() (string, error)
}

var (
	ID   = make(map[string]string)
	Logs = make(map[string]string)
)

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
		log.Printf("[genid] Error in generating ID for %s: %v\n", entityType, err)
		Logs[entityType] = fmt.Sprintf("ERROR: %v", err)
		return ""
	}

	log.Printf("[genid] ID for %s successfully created: %s\n", entityType, id)
	ID[entityType] = id
	Logs[id] = "OK"
	return id
}

func ShowAllID() {
	fmt.Printf("***---***\nAll users IDs:\n")
	num := 1
	for entityType, id := range ID {
		fmt.Printf("%d. Entity Type: %s, ID: %s\n", num, entityType, id)
		num += 1
	}

	fmt.Printf("***---***\nAll logs for ID:\n")
	num = 1
	for id, err := range Logs {
		fmt.Printf("%d. IDs: %s, Error: %s\n", num, id, err)
		num += 1
	}
	fmt.Println()
}
