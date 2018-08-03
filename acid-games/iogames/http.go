package iogames

import (
	"encoding/json"
	"io"
	"log"
)

//Response Custom Response Object
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var responses = map[string]string{
	"bad-validation":    `{"code": 400, "message": "A validation error has ocurred"}`,
	"no-db-connection":  `{"code": 451, "message": "Could not connect do database"}`,
	"insert-db-error":   `{"code": 500, "message": "Could not insert into the database"}`,
	"get-from-db-error": `{"code": 453, "message": "Could not get data from db"}`,
	"bad-encode-json":   `{"code": 454, "message": "Could not encode data to JSON"}`,
	"bad-json":          `{"code": 400, "message": "Invalid data format"}`,
	"sucess":            `{"code": 200, "message": "Operation succeed"}`,
	"internalError":     `{"code": 500, "message": "An internal problem has ocurred"}`,
	"pageNotFound":      `{"code": 404, "message": "Page not found"}`,
	"forbidden":         `{"code": 403, "message": "You don't have acess to this page"}`,
	"userNotFound":      `{"code": 404, "message": "Could not find user"}`,
}

//DecodeJSON .
func DecodeJSON(body io.Reader, entity interface{}) (err error) {
	decoder := json.NewDecoder(body)

	err = decoder.Decode(entity)

	if err != nil {
		log.Printf("erro ao decodificar json: %v", err)
		return
	}
	return
}

//EncodeJSON .
func EncodeJSON(entity interface{}) (entityEncoded []byte, err error) {
	entityToJSON, err := json.Marshal(entity)

	return entityToJSON, err
}
