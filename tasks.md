# Project Improvement Tasks

## 1. Integrate SQLite Instead of JSON Files
- +Set up a connection to SQLite.
- +Create a `notes` table to store note data.
- +Implement CRUD operations to work with SQLite.

## 2. Refactor Project Structure
- Separate project layers:
  - **Storage Layer (storage)**: Handles SQLite interactions.
  - **Service Layer (services)**: Business logic for handling notes.
  - **Handlers (handlers)**: Handles HTTP requests.

## 3. Use `gorilla/mux` for Routing
- +Integrate the `gorilla/mux` router.
- +Add routes with parameters (e.g., `/notes/{id}`).

## 4. Add Authentication and Authorization
- Implement user registration and login.
- Add token-based authentication (e.g., JWT).
- Restrict access to notes to their respective owners.

## 5. Input Validation
- Ensure that fields `title`, `content`, and `author` are properly validated.
- Set limits for the length of the title and content.

## 6. Improve Error Handling
- Create a centralized function for logging and returning errors.
- Implement custom error types with user-friendly messages.

## 7. Enhance RESTful API
- Add routes for filtering and sorting notes.
- Ensure proper HTTP status codes are returned for errors (e.g., `400`, `404`, `500`).

## 8. Swagger Documentation
- Document all endpoints using Swagger.
- Include request and response examples in the documentation.

## 9. Testing
- Write unit tests for CRUD operations.
- Implement integration tests for the API.

## 10. Deploy the Project
- Prepare deployment instructions for a production server.
- Consider using Docker for easy setup and deployment.

