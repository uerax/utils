package http

import (
	"encoding/json"
	"io"
	"net/http"
)

func Get(url string, t any) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}
	return json.Unmarshal(b, t)
}