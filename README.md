# Azion CertManager CLI

The **Azion CertManager CLI** is a command-line interface designed to simplify the management of SSL certificates on the Azion platform. This tool allows you to upload, update, list, and delete SSL certificates using your personal token.

## Getting Started

1. Make sure you have Go installed on your system.
2. Clone or download the project to your local machine.
3. Open a terminal and navigate to the project directory.

## Installation

No installation is required. You can run the Azion CertManager CLI directly from the source code.

## Usage

1. Run the CLI by executing the following command:

   ```bash
   go run main.go
   ```

2. The CLI will prompt you to enter your **Personal Token**. This token is required for authentication with the Azion platform.

3. Choose an operation from the list:
   - **1**: Upload a certificate to the Azion platform.
   - **2**: Update a certificate on the Azion platform.
   - **3**: List your certificates on the Azion platform.
   - **4**: Delete a certificate from the Azion platform.

4. Depending on your choice, follow the on-screen instructions.

5. After each operation, you will be asked if you want to perform another operation. Type **Y** to continue or **N** to exit the CLI.

## Features

- **Certificate Upload**: Upload a new SSL certificate to the Azion platform.
- **Certificate Update**: Update an existing SSL certificate on the Azion platform.
- **Certificate Listing**: List all certificates associated with your account on the Azion platform.
- **Certificate Deletion**: Delete a certificate from the Azion platform.

## Note

- This CLI is designed to interact with the Azion platform's certificate management services.
- Ensure you have a valid personal token from Azion before using the CLI.
- The provided code is a basic implementation and might need adjustments based on the actual implementation of the `inter` package.
- Always exercise caution while performing operations that modify or delete data.
- Make sure to follow best practices for securing and managing personal tokens.

## Credits

This CLI is developed by the community and is not an official Azion product.

## License

This project is licensed under the [MIT License](LICENSE).
