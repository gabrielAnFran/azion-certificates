package internal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// DeleteCertificate
//   - Receives a personal token.
//   - I/O operations.
//   - Calls and creates a request.
//   - Do the request.
//   - Returns an error, if it does not exist, returns nil.
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
	fmt.Println("Certificate: ", ID, " was succesfully deleted.")

	return nil
}
