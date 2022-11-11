package utils

import (
	"encoding/json"
	"fmt"
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

func ConvertToStruct(res json.RawMessage, customType interface{}) {
	json.Unmarshal(res, customType)
}
