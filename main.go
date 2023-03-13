package main

import (
	inter "certificates-azion-api/internal"
	"fmt"
)

func main() {
	var personalToken, process string
	continueProcess := true

	fmt.Println("Hey, there! Welcome to Azion CertManager")

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
			err := inter.CertificateCreateHandler(personalToken)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		if process == "2" {
			err := inter.CertificateUpdateHandler(&personalToken)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		if process == "3" {
			err := inter.CertificatesList(&personalToken)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		if process == "4" {
			err := inter.DeleteCertificate(&personalToken)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		resumeProcess := ""
		fmt.Println("Do you wish to do any other operation? Y, N")
		fmt.Scanf("%s", &resumeProcess)
		if resumeProcess == "Y" {
			continue
		} else {
			break
		}

	}
}
