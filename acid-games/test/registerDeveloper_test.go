package test

import (
	"Vitor-Bichara/servidorGO/iogames"
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func TestRegisterDeveloperSucess(t *testing.T) {
	expected := `{"code": 200, "message": "Operation succeed"}`

	devUniqueUsername := fmt.Sprintf("test_%s", iogames.RandomString(10))

	devS := `{
			"username": "%s",
			"password": "123456"
		}`

	dev := fmt.Sprintf(devS, devUniqueUsername)

	r := bytes.NewReader([]byte(dev))

	resp, _ := http.Post("http://"+iogames.SERVER_HOST+"/developers/register", "application/json", r)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	response := buf.String()

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}
}
func TestRegisterDeveloperBadJSONError(t *testing.T) {
	expected := `{"code": 400, "message": "Invalid data format"}`

	dev := `{
			"username": "test",
			"password": "123456"
			{}asdd
		}`

	r := bytes.NewReader([]byte(dev))

	resp, _ := http.Post("http://"+iogames.SERVER_HOST+"/developers/register", "application/json", r)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	response := buf.String()

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}
}

func TestRegisterDeveloperValidationErrorByMissingRequiredField(t *testing.T) {
	expected := `{"code": 400, "message": "A validation error has ocurred"}`
	dev := `{
			"username": "test"
		}`

	r := bytes.NewReader([]byte(dev))

	resp, _ := http.Post("http://"+iogames.SERVER_HOST+"/developers/register", "application/json", r)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	response := buf.String()

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}
}

func TestRegisterDeveloperValidationErrorByMaxLenght(t *testing.T) {
	expected := `{"code": 400, "message": "A validation error has ocurred"}`
	dev := `{
			"username": "test",
			"password": "123456789123456789123456789"
		}`

	r := bytes.NewReader([]byte(dev))

	resp, _ := http.Post("http://"+iogames.SERVER_HOST+"/developers/register", "application/json", r)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	response := buf.String()

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}
}
