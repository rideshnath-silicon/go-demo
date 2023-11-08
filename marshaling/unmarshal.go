package main

import (
	"encoding/json"
)

func unmarshal(s string, d Employee) Employee {

	data := []byte(s)

	unmData := json.Unmarshal(data, &d)

	if unmData != nil {
		panic("err" + unmData.Error())
	}

	return d
}
