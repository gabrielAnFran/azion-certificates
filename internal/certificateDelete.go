package internal

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func DeleteCertificate(token *string) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Inform the ID of the certificate you want to delete: ")
	ID, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	ID = strings.ReplaceAll(ID, "\n", "")

	// Reads the input

	r, err := Request("DELETE", ID, *token, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(r)

	if err != nil {
		return err
	}
	fmt.Println("Certificate: ", ID, " was succesfully deleted.")

	return nil
}
