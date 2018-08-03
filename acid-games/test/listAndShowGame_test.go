package test

import (
	"Vitor-Bichara/servidorGO/iogames"
	"net/http"
	"testing"
)

func TestListAllGamesSucess(t *testing.T) {
	expected := 200

	resp, _ := http.Get("http://" + iogames.SERVER_HOST + "/")

	if resp.StatusCode != expected {
		t.Errorf("The http response code was: %v Expected: %v", resp.StatusCode, expected)
	}
}

func TestShowGameSucess(t *testing.T) {
	expected := 200

	resp, _ := http.Get("http://" + iogames.SERVER_HOST + "/ih")

	if resp.StatusCode != expected {
		t.Errorf("The http response code was: %v Expected: %v", resp.StatusCode, expected)
	}
}

func TestShowGameFailByGameNotFound(t *testing.T) {
	expected := 404

	resp, _ := http.Get("http://" + iogames.SERVER_HOST + "/asdasdas")

	if resp.StatusCode != expected {
		t.Errorf("The http response code was: %v Expected: %v", resp.StatusCode, expected)
	}
}
