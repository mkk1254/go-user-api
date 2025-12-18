# Go User Management API

A RESTful API built using Go (Golang) to manage users and dynamically calculate their age based on date of birth. This project was developed as part of the Ainyx Solutions Software Engineering Intern backend assignment.

---

## Tech Stack

- Go (Golang)
- GoFiber
- PostgreSQL (Supabase)
- SQLC
- go-playground/validator
- Uber Zap Logger

---

## Features

- Create a user with name and date of birth
- Fetch a user by ID with dynamically calculated age
- List all users with age
- Update user details
- Delete user
- Clean RESTful API design
- Age is calculated dynamically (not stored in database)

---

## Project Structure

.
├── cmd/
│   └── server/
│       └── main.go
├── config/
│   └── config.go
├── db/
│   ├── migrations/
│   └── sqlc/
│       ├── queries/
│       │   └── users.sql
│       ├── db.go
│       ├── models.go
│       └── users.sql.go
├── internal/
│   ├── handler/
│   │   └── user_handler.go
│   ├── logger/
│   │   └── logger.go
│   ├── middleware/
│   ├── models/
│   │   └── user.go
│   ├── routes/
│   │   └── routes.go
│   └── service/
│       └── age.go
├── sqlc.yaml
├── go.mod
├── go.sum
└── README.md

---

## Database Schema

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

---

## API Endpoints

### Create User
POST /users

Request Body:
{
  "name": "Alice",
  "dob": "1990-05-10"
}

---

### Get User by ID
GET /users/{id}

Response:
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 35
}

---

### List All Users
GET /users

Response:
[
  {
    "id": 1,
    "name": "Alice",
    "dob": "1990-05-10",
    "age": 35
  }
]

---

### Update User
PUT /users/{id}

Request Body:
{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}

---

### Delete User
DELETE /users/{id}

Response:
204 No Content

---

## Setup and Run

1. Clone the repository

git clone <your-repository-url>
cd go-user-api

2. Set environment variable

DATABASE_URL=postgresql://<username>:<password>@<pooler-host>:<port>/postgres?sslmode=require

Example:
postgresql://postgres.projectref:password@aws-x-region.pooler.supabase.com:5432/postgres?sslmode=require

3. Run the server

go run cmd/server/main.go

Server will start at:
http://localhost:3000

---

## Testing

The API can be tested using PowerShell, curl, or any REST client.

Example:
irm http://localhost:3000/users

---

## Design Decisions

- SQLC was used for type-safe database access
- Age is calculated dynamically using Go time package
- Supabase used as managed PostgreSQL database
- Layered architecture improves readability and maintainability
- Validation handled using go-playground/validator
- Logging implemented using Uber Zap

---

## Notes

- Authentication was not required and hence not implemented
- Pagination was optional and not included
- Focus was on correctness, clarity, and clean structure

---

## Author

Monal Khatri  

