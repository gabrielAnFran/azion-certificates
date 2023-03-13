package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type CertificateListResponse struct {
	Results []Certificate `json:"results"`
}

type Certificate struct {
	ID              int           `json:"id"`
	Name            string        `json:"name"`
	Issuer          string        `json:"issuer"`
	SubjectName     []interface{} `json:"subject_name"`
	Validity        string        `json:"validity"`
	Status          string        `json:"status"`
	CertificateType string        `json:"certificate_type"`
}

// CertificatesList
//   - Receives a personal token.
//   - I/O operations.
//   - Calls and creates a request.
//   - Do the request.
//   - Returns an error, if it does not exist, returns nil.
func CertificatesList(token *string) error {

	r, err := Request("GET", "", *token, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(r)

	if err != nil {
		bytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\nError")
		fmt.Println(string(bytes))
		return err
	}
	b, _ := io.ReadAll(res.Body)
	var response CertificateListResponse
	json.Unmarshal(b, &response)

	for _, cert := range response.Results {
		fmt.Println("ID: ", cert.ID, " - Name: ", cert.Name, " - Status: ", cert.Status)
	}
	return nil
}
