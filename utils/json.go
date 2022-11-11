package utils

import (
	"encoding/json"
	"fmt"
	"github.com/BzingaApp/user-svc/entities"
	"reflect"
)

func IsEqualJSON(s1, s2 string) (flag bool, err error) {
	var o1 interface{}
	var o2 interface{}

	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		err = fmt.Errorf("error mashalling string 1 :: %s", err.Error())
		return
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		err = fmt.Errorf("error mashalling string 2 :: %s", err.Error())
		return
	}
	flag = reflect.DeepEqual(o1, o2)
	return
}

func IsEqualGenericResponse(s1, s2 string) (flag bool, err error) {
	var g1, g2 entities.GenericResponse

	err = json.Unmarshal([]byte(s1), &g1)
	if err != nil {
		err = fmt.Errorf("error mashalling string 1 :: %s", err.Error())
		return
	}
	err = json.Unmarshal([]byte(s2), &g2)
	if err != nil {
		err = fmt.Errorf("error mashalling string 2 :: %s", err.Error())
		return
	}
	g1.Data = ""
	g2.Data = ""
	flag = reflect.DeepEqual(g1, g2)
	return
}
