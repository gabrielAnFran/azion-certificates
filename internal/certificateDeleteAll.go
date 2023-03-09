package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func DeleteAllCertificates(token *string) error {
	r, err := Request("GET", "", *token, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(r)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	b, _ := io.ReadAll(res.Body)

	var responseList CertificateListResponse
	json.Unmarshal(b, &responseList)

	for _, cert := range responseList.Results {
		r, err := Request("DELETE", strconv.Itoa(cert.ID), *token, nil)
		if err != nil {
			return err
		}

		client := &http.Client{}
		res, err := client.Do(r)
		if err != nil {
			return err
		}

		fmt.Println(res.Status)
	}

	return nil
}
