package internal

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type CertificateCreateResponseResult struct {
	Result CertificateCreateResponse `json:"results"`
}

type CertificateCreateResponse struct {
	ID              int           `json:"id"`
	Name            string        `json:"name"`
	Issuer          string        `json:"issuer"`
	SubjectName     []interface{} `json:"subject_name"`
	Validity        string        `json:"validity"`
	Status          string        `json:"status"`
	CertificateType string        `json:"certificate_type"`
}

type CertificateBody struct {
	Name        string `json:"name"`
	Certificate string `json:"certificate"`
	Key         string `json:"private_key"`
}

func CertificateCreateHandler(personalToken string) error {
	var cert, privateKey, certName string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Plase, provide your certificate's name: ")

	// Reads the input
	certName, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	certName = strings.Replace(certName, "\n", "", -1)

	fmt.Println("Plase, provide the path to your certificate: ")

	// Reads the input
	// Replaces \n
	cert, err = reader.ReadString('\n')
	if err != nil {
		return err
	}
	cert = strings.Replace(cert, "\n", "", -1)
	// Open the file
	datCert, err := os.ReadFile(cert)
	if err != nil {
		return err
	}

	fmt.Println("Plase, provide the path to your certificate's private key: ")
	privateKey, err = reader.ReadString('\n')
	if err != nil {
		return err
	}
	privateKey = strings.Replace(privateKey, "\n", "", -1)
	dataKey, err := os.ReadFile(privateKey)
	if err != nil {
		return err
	}

	privateKeyCertificate := string(dataKey)
	certificate := string(datCert)

	certificate = strings.ReplaceAll(certificate, "\t", "")
	privateKeyCertificate = strings.ReplaceAll(privateKeyCertificate, "\t", "")

	err = NewCertificate(&personalToken, &certName, &certificate, &privateKeyCertificate)
	if err != nil {
		return err
	}
	return nil
}

func NewCertificate(token, name, certificate, priv_key *string) error {
	body := CertificateBody{
		Name:        *name,
		Certificate: *certificate,
		Key:         *priv_key,
	}

	bodyencode, err := json.Marshal(body)
	if err != nil {
		return err
	}

	r, err := Request("POST", "", *token, bytes.NewReader(bodyencode))
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(r)

	if err != nil {
		return err
	}

	b, _ := io.ReadAll(res.Body)
	var response CertificateCreateResponseResult

	err = json.Unmarshal(b, &response)
	if err != nil {
		return err
	}

	fmt.Println("Certificated created succesfully")
	fmt.Println("ID: ", response.Result.ID)
	fmt.Println("Name: ", response.Result.Name)
	fmt.Println("Status: ", response.Result.Status)

	return nil
}
