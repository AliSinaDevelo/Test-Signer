# Test-Signer
 a service that accepts a set of answers and questions and signs that the user has finished the " test " at this point in time

## Requirements

- Go 
- PostgreSQL

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/AliSinaDevelo/Test-Signer

2. Set up PostgreSQl:
Create a database named test_signer.

3. Set up Environment variables:
Create a .env file in the root directory with the following content:(example)

    ```plaintext
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=your_username
    DB_PASSWORD=your_password
    DB_NAME=test_signer

## Running the Application
1. Navigate to the project dir:
   ```bash
   cd test-signer

2. Run the test Signer service:
   ```bash
   go run cmd/Test-Signer/main.go

3. Accessing endpoints:
/sign: POST request to sign answers and questions.
/verify: GET request to verify a signature.

## Testing Endpoints
se tools like curl, Postman, or a web browser to test the /sign and /verify endpoints.

## Database Verification
Use pgAdmin or another PostgreSQL client to verify stored signatures.

## Cleanup
To stop the server, press Control + C in the terminal running the Test Signer service.

## Aditional Notes
Adjust environment variables and configurations based on your environment setup.
Ensure proper database permissions and security practices are followed.
Consider adding tests for improved code reliability and maintainability.