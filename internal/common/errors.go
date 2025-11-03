package common

import (
	"fmt"
	"log"
)

var Logs = make(map[string]string)

func LogError(source string, err error) {
	if err != nil {
		msg := fmt.Sprintf("[%s] %v", source, err)
		log.Println(msg)
		Logs[source] = msg
	}
}

func ShowLogsErr() {
	fmt.Printf("***---***\nAll logs for ID:\n")
	num := 1
	for id, err := range Logs {
		fmt.Printf("%d. IDs: %s, Error: %s\n", num, id, err)
		num += 1
	}
}
