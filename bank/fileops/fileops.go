package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteValueToFile(value float64, filename string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueText), 0644)
}

func ReadValueFromFile(filename string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return defaultValue, errors.New("couldn't read the file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return defaultValue, errors.New("couldn't parse the file")
	}

	return value, nil
}