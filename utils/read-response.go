package utils

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func ReadResonse(resp *http.Response) (res json.RawMessage, err error) {
	res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("failed to read the response -%v", res)
	}
	return
}

func ConvertJSONToGoType(res json.RawMessage, customType interface{}) {
	err := json.Unmarshal(res, customType)

	if err != nil {
		log.Error("failed to unmarshall due to ", err.Error())
	}
}
