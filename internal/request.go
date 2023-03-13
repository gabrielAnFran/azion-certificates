package internal

import (
	"io"
	"net/http"
)

const (
	certificateURL = "https://api.azionapi.net/digital_certificates"
)

// Request
//
//	Receives:
//	* Http method
//	* ID, if it is an update, delete or get by id it is a required value
//	* Azion Personal Token
//	* Request body
//	Returns:
//	* Request and an error, if erros does not exist, returns nil
func Request(method, ID, token string, body io.Reader) (*http.Request, error) {

	if ID != "" {
		ID = "/" + ID
	}
	r, err := http.NewRequest(method, certificateURL+ID, body)
	if err != nil {
		return nil, err
	}

	// Required headers
	r.Header.Add("Accept", "application/json;version=3")
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Token "+token)
	r.Header.Add("User-Agent", "Azion_CLI/development")

	return r, nil

}
