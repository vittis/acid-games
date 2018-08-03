package test

import (
	"Vitor-Bichara/servidorGO/iogames"
	"bytes"
	"net/http"
	"testing"
)

func TestUpvoteGameSucess(t *testing.T) {
	expected := 200

	resp, _ := http.Post("http://"+iogames.SERVER_HOST+"/ih/upvote", "", nil)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if resp.StatusCode != expected {
		t.Errorf("The http response code was: %v Expected: %v", resp.StatusCode, expected)
	}
}
func TestUDownvoteGameSucess(t *testing.T) {
	expected := 200

	resp, _ := http.Post("http://"+iogames.SERVER_HOST+"/ih/downvote", "", nil)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if resp.StatusCode != expected {
		t.Errorf("The http response code was: %v Expected: %v", resp.StatusCode, expected)
	}
}
