package middleware

import (
	"fmt"
	"strconv"
	"time"
)

func errorCode(expected int, received int) string {
	return fmt.Sprintf("Expecting Http Status Code: %v, but received: %v", expected, received)
}

func delay() {
	time.Sleep(1 * time.Second)
}

var port = 8000

func getAvaliablePort() string {
	currentPort := port
	port++
	return strconv.Itoa(currentPort)
}
