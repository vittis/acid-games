package test

import (
	"Vitor-Bichara/servidorGO/iogames"
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestShowSignUpFormSucess(t *testing.T) {
	//expectedError := `{"code": 500, "message": "An internal problem has ocurred"}`

	expected, _ := ioutil.ReadFile("../iogames/templates/sign_in_dev.html")

	expectedStringFormat := string(expected)

	resp, _ := http.Get("http://" + iogames.SERVER_HOST + "/developers/sign_in")

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	response := buf.String()

	if strings.Compare(response, expectedStringFormat) == 0 {
		t.Errorf("The http response was: %v Expected: %v", response, expectedStringFormat)
	}
}
