# Simple Social Media Backend Project

This is a simple social media backend project built with Golang. It includes authentication using JWT (JSON Web Tokens) and bcrypt for password hashing. The project also incorporates hooks for `beforeCreate` and `beforeUpdate` operations. The main functionalities provided by this backend are CRUD (Create, Read, Update, Delete) operations for photos, social media posts, users, and comments.

## Features

- User authentication using JWT (JSON Web Tokens)
- Password hashing using bcrypt
- Hooks for `beforeCreate` and `beforeUpdate` operations
- CRUD operations for photos, social media posts, users, and comments

## Prerequisites

Before running the project, ensure you have the following installed:

- Golang (version X.X.X or higher)
- PostgreSQL (version X.X.X or higher)

## Getting Started

1. Clone the repository to your local machine:

```
git clone https://github.com/your-username/simple-social-media-backend.git
```

2. Navigate to the project directory:

```
cd simple-social-media-backend
```

3. Install the dependencies:

```
go mod download
```

4. Create a PostgreSQL database for the project.

5. Update the database configuration in the `config/config.go` file to match your database settings:

```go
const (
    DBHost     = "localhost"
    DBPort     = 5432
    DBUser     = "your-database-username"
    DBPassword = "your-database-password"
    DBName     = "your-database-name"
)
```

6. Run the database migrations to create the necessary tables:

```
go run migration/migration.go
```

7. Start the server:

```
go run main.go
```

The server should now be running on `http://localhost:8000`.

## Endpoints

The following endpoints are available for interacting with the API:

### Authentication

- `POST /api/auth/signup`: Register a new user.
- `POST /api/auth/login`: Log in an existing user.

### Photos

- `GET /api/photos`: Retrieve all photos.
- `GET /api/photos/{id}`: Retrieve a specific photo.
- `POST /api/photos`: Create a new photo.
- `PUT /api/photos/{id}`: Update a photo.
- `DELETE /api/photos/{id}`: Delete a photo.

### Social Media Posts

- `GET /api/posts`: Retrieve all social media posts.
- `GET /api/posts/{id}`: Retrieve a specific social media post.
- `POST /api/posts`: Create a new social media post.
- `PUT /api/posts/{id}`: Update a social media post.
- `DELETE /api/posts/{id}`: Delete a social media post.

### Users

- `GET /api/users`: Retrieve all users.
- `GET /api/users/{id}`: Retrieve a specific user.
- `PUT /api/users/{id}`: Update a user.
- `DELETE /api/users/{id}`: Delete a user.

### Comments

- `GET /api/comments`: Retrieve all comments.
- `GET /api/comments/{id}`: Retrieve a specific comment.
- `POST /api/comments`: Create a new comment.
- `PUT /api/comments/{id}`: Update a comment.
- `DELETE /api/comments/{id}`: Delete a comment.

## Contributing

If you want to contribute to this project, you can follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them with descriptive messages.
4. Push your changes to your forked repository.
5. Submit a pull request to the original
