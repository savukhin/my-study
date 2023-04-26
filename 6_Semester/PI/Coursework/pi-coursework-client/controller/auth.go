package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func Auth(username, password string) error {
	request := map[string]string{"username": username, "password": password}
	jsonSring, _ := json.Marshal(request)
	r := bytes.NewReader(jsonSring)
	resp, err := http.Post("http://example.com/upload", "application/json", r)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {
		return nil
	}
	return errors.New("failed to auth")
}
