package gohelix

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

func (h *Helix) Request(method string, url string, body []byte, headers map[string]string) ([]byte, error) {
	// Reset HttpClient
	h.HttpClient = http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	res, err := h.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)
	reqBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return reqBody, nil
}
