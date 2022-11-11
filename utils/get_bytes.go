package utils

import (
	"bytes"
	"encoding/gob"
	log "github.com/sirupsen/logrus"
)

func GetBytes(req interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(req)
	if err != nil {
		log.Debug("get-bytes ", err)
		return nil, err
	}
	return buf.Bytes(), nil
}
