# task-5-pbi-btpns-MarsyaPutra
 
# MOBILEBANKINGAPI

MOBILEBANKINGAPI is a GoLang-based API for a mobile banking application that focuses on improving user engagement. It includes features such as user registration, login, updating user information, deleting user accounts, and managing profile pictures.

## Table of Contents
- [Installation](#installation)
- [Project Structure](#project-structure)
- [Configuration](#configuration)
- [Running the API](#running-the-api)
- [API Endpoints](#api-endpoints)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/your-username/MOBILEBANKINGAPI.git
   cd MOBILEBANKINGAPI
   
2. **Database Setup:**

Create a MySQL database for the application.
Update the database connection details in SQLdatabase/database.go.

3. **Dependencies:**

Install the Go dependencies using the following command:

go mod tidy

4. **Initialize the Database:**

Run the following command to initialize the database tables:
go run main.go initdb

## Project Structure
The project follows the following directory structure:

MOBILEBANKINGAPI/
|-- app/
|-- controllers/
|-- helpers/
|-- models/
|-- router/
|-- SQLdatabase/
|-- uploads/
|-- main.go
|-- go.mod
|-- go.sum
app/: Application initialization logic.
controllers/: Controllers for handling various API endpoints.
helpers/: Helper functions used across the project.
models/: Data models for users and photos.
router/: API route definitions.
SQLdatabase/: Database initialization and connection logic.
uploads/: Directory to store uploaded photos.

## Configuration

Database configuration is in SQLdatabase/database.go.
JWT secret key is in helpers/token_helper.go.

## Running the API
Run the following command to start the API:

go run main.go
The API will be accessible at http://localhost:8080.

## API Endpoints
User Registration: POST /register
User Login: POST /login
Update User Information: PUT /user
Delete User Account: DELETE /user
Upload Photo: POST /photo
Retrieve Photos: GET /photo
Update Photo Information: PUT /photo/:id
Delete Photo: DELETE /photo/:id
Usage
Register a user, login, and obtain a JWT token.
Use the token for authorization when updating user information or managing photos.

## Contributing
Feel free to contribute to the project. Open an issue for discussions or submit a pull request with improvements.
