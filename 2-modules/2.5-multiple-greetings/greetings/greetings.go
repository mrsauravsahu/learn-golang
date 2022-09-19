package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	messageFormat := getRandomGreetFormat()
	message := fmt.Sprintf(messageFormat, name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRandomGreetFormat() string {
	greetFormats := []string{
		"Hi, %v! Welcome!",
		"Great to see you, %v!",
		"Bonjour, %v!",
		"Salut, %v",
		"Comment ça va, %v?",
	}

	return greetFormats[rand.Intn(len(greetFormats))]
}
