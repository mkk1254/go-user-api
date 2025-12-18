## Reasoning & Design Decisions

### Project Structure
The project was designed following the requested layered architecture:
- `handler` for HTTP request handling
- `service` for business logic
- `repository` for database access (SQLC)
- `routes` for route registration
- `config` and `logger` for setup and observability

Due to GitHub UI limitations while uploading an empty repository, the folder hierarchy appears flattened in the repository view.  
However, the logical separation is preserved through Go packages and imports, and the application structure matches the intended design.

### Database & SQLC
PostgreSQL is used as the database, with SQLC generating type-safe query code.  
The date of birth is stored in the database, while age is calculated dynamically using Go’s `time` package.

### Validation & Logging
Request validation is implemented using `go-playground/validator`, and logging is handled via Uber Zap for structured logs.

### API Design
RESTful principles are followed with clean HTTP status codes and clear request/response formats.
