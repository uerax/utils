package http

import (
	"encoding/json"
	"io"
	"net/http"
)

func Get(url string, parameter map[string]string, header map[string]string, t any) error {
	
	if len(parameter) != 0 {
		if url[len(url) - 1] != '?' {
			url += "?"
		}
		for k, v := range parameter {
			if url[len(url) - 1] != '?' {
				url += "&"
			}
			url += k + "=" + v
		}
	}

	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	for k, v := range header {
		r.Header.Set(k, v)
	}

	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}
	return json.Unmarshal(b, t)
}



func Post(url string, parameter io.Reader, header map[string]string, t any) error {
	req, err := http.NewRequest(http.MethodPost, url, parameter)
	if err != nil {
		return err
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &t)

}