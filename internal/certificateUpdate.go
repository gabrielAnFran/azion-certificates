package internal

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

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
	_, err = client.Do(r)

	if err != nil {
		return err
	}

	// err = CertificatesList(token)
	// if err != nil {
	// 	return err
	// }
	fmt.Println("Certificated Updated succesfully")
	return nil
}
