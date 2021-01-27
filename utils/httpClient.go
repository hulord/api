package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type HttpConPool struct {
	Conn *http.Client
}

var HttpClient *HttpConPool

func (h *HttpConPool) Request(url string, method string, data string, header map[string]string) (interface{}, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, err
	}
	header["User-Agent"] = "my browser"
	for h, v := range header {
		req.Header.Set(h, v)
	}
	response, err := h.Conn.Do(req)

	if err != nil {
		return nil, err
	} else if response != nil {
		defer response.Body.Close()

		r_body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		} else {
			return string(r_body), nil
		}
	} else {
		return nil, nil
	}
}


func Get(url string)(string, error){
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return string(body),err
}

