package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	// If no name was given, return an error
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people with a greeting message
func Hellos(names []string) (map[string]string, error) {
	// a map to associate names iwth messages.
	messages := make(map[string]string)

	// loop through the slice of names, calling the Hello function to get
	// a message each time
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved messsage with the name
		messages[name] = message
	}

	return messages, nil
}

// sets intial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
