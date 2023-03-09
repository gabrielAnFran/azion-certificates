package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
)

const (
	certificateURL = "https://api.azionapi.net/digital_certificates"
)

type CertificateBody struct {
	Name        string `json:"name"`
	Certificate string `json:"certificate"`
	Key         string `json:"private_key"`
}

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

func Request(method, ID, token string, body io.Reader) (*http.Request, error) {

	if ID != "" {
		ID = "/" + ID
	}
	spew.Dump(certificateURL + ID)
	r, err := http.NewRequest(method, certificateURL+ID, body)
	if err != nil {
		return nil, err
	}

	r.Header.Add("Accept", "application/json;version=3")
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Token "+token)
	r.Header.Add("User-Agent", "Azion_CLI/development")

	return r, nil

}

func main() {
	var personalToken, process string
	continueProcess := true

	fmt.Println("Hey, there! Welcome to Azion xxxxxxx")

	fmt.Println("How can I help you?")

	for continueProcess {
		fmt.Println("Enter 1 - Upload a certificate to the Azion platform")
		fmt.Println("Enter 2 - Update a certificate to the Azion platform")
		fmt.Println("Enter 3 - List your certificates on the Azion platform")
		fmt.Println("Enter 4 - Delete a certificate to the Azion platform")

		fmt.Scanf("%s", &process)

		fmt.Println("Plase, provide your Personal Token: ")
		fmt.Scanf("%s", &personalToken)

		// If the user wants to upload a new certificate
		if process == "1" {
			err := CertificateCreateHandler(personalToken)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		if process == "2" {
			err := CertificateUpdateHandler(&personalToken)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		if process == "3" {
			err := CertificatesList(&personalToken)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		if process == "4" {
			err := DeleteAllCertificates(&personalToken)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		resumeProcess := ""
		fmt.Println("Wish to do some other operatios? Y, N")
		fmt.Scanf("%s", &resumeProcess)
		if resumeProcess == "Y" {
			continue
		} else {
			break
		}

	}
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
		spew.Dump(res.Body)
		spew.Dump(err.Error())
		return err
	}

	b, _ := io.ReadAll(res.Body)
	var response CertificateCreateResponseResult

	err = json.Unmarshal(b, &response)
	if err != nil {
		return err
	}

	spew.Dump(response)
	fmt.Println("Certificated created succesfully")
	fmt.Println("ID: ", response.Result.ID)
	fmt.Println("Name: ", response.Result.Name)
	fmt.Println("Status: ", response.Result.Status)

	return nil
}

func CertificatesList(token *string) error {

	r, err := Request("GET", "", *token, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(r)

	if err != nil {
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

func CertificateUpdateHandler(personalToken *string) error {
	var cert, privateKey, certName, id string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Your digital certificates: ")
	CertificatesList(personalToken)
	time.Sleep(5)

	fmt.Println("Plase, provide your certificate's ID: ")

	// Reads the input
	id, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	id = strings.Replace(id, "\n", "", -1)

	fmt.Println("Plase, provide your certificate's name: ")

	// Reads the input
	certName, err = reader.ReadString('\n')
	if err != nil {
		return err
	}

	certName = strings.Replace(certName, "\n", "", -1)
	time.Sleep(10)

	fmt.Println("Plase, provide the path to your certificate: ")

	// Reads the input
	// Replaces \n
	cert, err = reader.ReadString('\n')
	if err != nil {
		return err
	}
	cert = "/home/franz/Downloads/A_cert.pem"
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
	privateKey = "/home/franz/Downloads/A_key.pem"
	privateKey = strings.Replace(privateKey, "\n", "", -1)
	dataKey, err := os.ReadFile(privateKey)
	if err != nil {
		return err
	}

	privateKeyCertificate := string(dataKey)
	certificate := string(datCert)

	certificate = strings.ReplaceAll(certificate, "\t", "")
	privateKeyCertificate = strings.ReplaceAll(privateKeyCertificate, "\t", "")

	err = CertificateUpdate(personalToken, &certName, &certificate, &privateKeyCertificate, &id)
	if err != nil {
		return err
	}
	return nil
}

func CertificateUpdate(token, name, certificate, priv_key, id *string) error {
	body := CertificateBody{
		Name:        *name,
		Certificate: *certificate,
		Key:         *priv_key,
	}

	bodyencode, err := json.Marshal(body)
	if err != nil {
		return err
	}

	r, err := Request("PUT", *id, *token, bytes.NewReader(bodyencode))
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(r)

	if err != nil {
		spew.Dump(res.Body)
		spew.Dump(err.Error())
		return err
	}

	// err = CertificatesList(token)
	// if err != nil {
	// 	return err
	// }
	b, _ := io.ReadAll(res.Body)
	spew.Dump(string(b))
	spew.Dump("Certificated Updated succesfully")
	return nil
}

func DeleteAllCertificates(token *string) error {
	r, err := Request("GET", "", *token, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(r)

	if err != nil {
		spew.Dump(res.Body)
		spew.Dump(err.Error())
		return err
	}
	b, _ := io.ReadAll(res.Body)

	var responseList CertificateListResponse
	json.Unmarshal(b, &responseList)

	spew.Dump(responseList)

	for _, cert := range responseList.Results {
		r, err := Request("DELETE", strconv.Itoa(cert.ID), *token, nil)
		if err != nil {
			return err
		}

		client := &http.Client{}
		res, err := client.Do(r)
		if err != nil {
			spew.Dump(res.Body)
			spew.Dump(err.Error())
			return err
		}

		spew.Dump(res.Status)
	}

	return nil
}
