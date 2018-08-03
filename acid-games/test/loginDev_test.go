package test

import (
	"Vitor-Bichara/servidorGO/iogames"
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func TestLoginDeveloperSucess(t *testing.T) {
	expected := `{"code": 200, "message": "Operation succeed"}`

	devUniqueUsername := fmt.Sprintf("test_%s", iogames.RandomString(10))

	devS := `{
			"username": "teste",
			"password": "123"
		}`

	dev := fmt.Sprintf(devS, devUniqueUsername)

	r := bytes.NewReader([]byte(dev))

	resp, _ := http.Post("http://"+iogames.SERVER_HOST+"/developers/login", "application/json", r)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	response := buf.String()

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}
}

func TestLoginDeveloperFailByNoExistingUser(t *testing.T) {
	expected := `{"code": 404, "message": "Could not find user"}`

	devUniqueUsername := fmt.Sprintf("test_%s", iogames.RandomString(10))

	devS := `{
			"username": "teste123123",
			"password": "123123123"
		}`

	dev := fmt.Sprintf(devS, devUniqueUsername)

	r := bytes.NewReader([]byte(dev))

	resp, _ := http.Post("http://"+iogames.SERVER_HOST+"/developers/login", "application/json", r)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	response := buf.String()

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}
}
func TestLoginDeveloperFailByBadJson(t *testing.T) {
	expected := `{"code": 400, "message": "Invalid data format"}`

	devUniqueUsername := fmt.Sprintf("test_%s", iogames.RandomString(10))

	devS := `{
			"useasdadasdrname": "teste123123",
			"password": "123123123"asda
			sdasd
		}`

	dev := fmt.Sprintf(devS, devUniqueUsername)

	r := bytes.NewReader([]byte(dev))

	resp, _ := http.Post("http://"+iogames.SERVER_HOST+"/developers/login", "application/json", r)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	response := buf.String()

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}
}
