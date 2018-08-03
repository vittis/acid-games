package test

import (
	"Vitor-Bichara/servidorGO/iogames"
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func TestRegisterGameSucess(t *testing.T) {
	expected := `{"code": 200, "message": "Operation succeed"}`

	gameUniqueTitle := fmt.Sprintf("test_%s", iogames.RandomString(10))

	gameS := `{
				"title": "%s",
				"url": "http://www.google.com",
				"description": "oloco",
				"category": "rpg",
				"thumbnail": "asdasdads",
				"owner": "fsssag"
			}`

	game := fmt.Sprintf(gameS, gameUniqueTitle)

	r := bytes.NewReader([]byte(game))

	resp, _ := http.Post("http://"+iogames.SERVER_HOST+"/developer/games/new_game", "application/json", r)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	response := buf.String()

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}
}

func TestRegisterGameBadJSONError(t *testing.T) {
	expected := `{"code": 400, "message": "Invalid data format"}`
	game := `{
				"title": "karpov",
				"url": "http://www.google.com",
				"description": "oloco",
				"category": "rpg",
				"thumbnail": "asdasdads",
				"owner": "fsssag",
				{}
					asdasd
			}`

	r := bytes.NewReader([]byte(game))

	resp, _ := http.Post("http://"+iogames.SERVER_HOST+"/developer/games/new_game", "application/json", r)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	response := buf.String()

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}
}

func TestRegisterGameValidationErrorByWrongFormatURL(t *testing.T) {
	expected := `{"code": 400, "message": "A validation error has ocurred"}`
	game := `{
				"title": "karpov",
				"url": "bad-url-format",
				"description": "oloco",
				"category": "rpg",
				"thumbnail": "asdasdads",
				"owner": "fsssag"
			}`

	r := bytes.NewReader([]byte(game))

	resp, _ := http.Post("http://"+iogames.SERVER_HOST+"/developer/games/new_game", "application/json", r)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	response := buf.String()

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}
}

func TestRegisterGameValidationErrorByNoExistingCategory(t *testing.T) {
	expected := `{"code": 400, "message": "A validation error has ocurred"}`
	game := `{
				"title": "karpov",
				"url": "http://www.google.com",
				"description": "oloco",
				"category": "wasdl",
				"thumbnail": "asdasdads",
				"owner": "fsssag"
			}`

	r := bytes.NewReader([]byte(game))

	resp, _ := http.Post("http://"+iogames.SERVER_HOST+"/developer/games/new_game", "application/json", r)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	response := buf.String()

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}
}

func TestRegisterGameValidationErrorByMissingRequiredField(t *testing.T) {
	expected := `{"code": 400, "message": "A validation error has ocurred"}`
	game := `{
				"title": "karpov",
				"url": "http://www.google.com",
				"description": "oloco",
				"category": "bad-category",
				"thumbnail": "asdasdads"
			}`

	r := bytes.NewReader([]byte(game))

	resp, _ := http.Post("http://"+iogames.SERVER_HOST+"/developer/games/new_game", "application/json", r)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	response := buf.String()

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}
}

func TestRegisterGameValidationErrorByMaxLenght(t *testing.T) {
	expected := `{"code": 400, "message": "A validation error has ocurred"}`
	game := `{
				"title": "karpovasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasdasd",
				"url": "http://www.google.com",
				"description": "oloco",
				"category": "rpg",
				"thumbnail": "asdasdads",
				"owner": "fsssag"
			}`

	r := bytes.NewReader([]byte(game))

	resp, _ := http.Post("http://"+iogames.SERVER_HOST+"/developer/games/new_game", "application/json", r)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	response := buf.String()

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}
}
