package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
	"strconv"
	"time"
)

// ConvertStringIntoFloat ... to convert string into float.
func ConvertStringIntoFloat(numeric string) float64 {
	if convertedNumber, err := strconv.ParseFloat(numeric, 64); err == nil {
		return convertedNumber
	}
	return 0
}

// ConvertNumericIntoString ... convert numeric(integer,float) into string.
func ConvertNumericIntoString[constraint int | int64 | float64](number interface{}) string {
	switch number.(type) {
	case int:
		return strconv.Itoa(number.(int))
	case int64:
		return strconv.Itoa(int(number.(int64)))
	case float64:
		return fmt.Sprintf("%.2f", number.(float64))
	}
	return ""
}

func ConvertStringIntoBool(data string) bool {
	if boolean, err := strconv.ParseBool(data); err != nil {
		return boolean
	}
	log.Println("error converting string into bool value ")
	return false
}

// ConvertStringIntoInt ... convert string into integer.
func ConvertStringIntoInt(numeric string) int {
	convertedNumber, err := strconv.Atoi(numeric)
	if err != nil {
		log.Println("error converting string into integer value ")
		return 0
	} else {
		return convertedNumber
	}
}
func ConvertFloat64toTimeInHoursMin(f float64) float32 {
	t := int(math.Round(f))
	timeout := time.Duration(t)
	login := (timeout * time.Minute).Hours()
	output := math.Pow(10, float64(2))
	return float32(math.Round(login*output) / output)
}
