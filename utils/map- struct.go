package utils

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// map[string]interface{}
//
//goland:noinspection ALL
func ConvertMapToAny(req interface{}, wres interface{}) (err error) {
	jsonStr, err := json.Marshal(req)
	if err != nil {
		return err
	}
	// Convert json string to struct
	if err := json.Unmarshal(jsonStr, wres); err != nil {
		log.Error(err)
		return err
	}
	return
}
