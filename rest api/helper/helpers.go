package helper

import (
	"encoding/json"
	"io"
	"net/http"
)


func GetMap[T any](s string, structs T) T {
	data := []byte(s)

	err := json.Unmarshal(data, &structs)

	if err != nil {
		panic("Error: " + err.Error())
	}

	return structs
}

func RequestBody[T any](w http.ResponseWriter, r *http.Request, Struct T) T {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		panic("Error: " + err.Error())
	}
	defer r.Body.Close()
	requestData := string(body)
	var bodyData T
	bodyData = GetMap(requestData, bodyData)
	return bodyData
}

func ApiSuccess(w http.ResponseWriter, data interface{}, MessageCode int) {
	type Responses struct {
		Message string
		Success int
		Data    interface{}
	}
	w.Header().Set("Content-Type", "Application/json")
	messages := messages(MessageCode)
	Response := Responses{Message: messages, Success: 1, Data: data}
	err := json.NewEncoder(w).Encode(Response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ApiFailure(w http.ResponseWriter, MessageCode int) {
	type Responses struct {
		Message string
		Success int
	}
	w.Header().Set("Content-Type", "Application/json")
	messages := messages(MessageCode)
	Response := Responses{Message: messages, Success: 0}
	err := json.NewEncoder(w).Encode(Response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

