# Golang Task - REST API for User Authentication and Data Storage

This is a Golang-based REST API project that provides user authentication using email and password, and allows authenticated users to store their personal details. The API uses JWT (JSON Web Token) for secure authentication.

## Features

1. **User Authentication** (`/login`):
   - Endpoint to authenticate users using their email and password.
   - Returns a JWT token upon successful authentication.

2. **User Registration** (`/signup`):
   - Endpoint to create a new user account.
   - Stores user details like name, email, and phone in the database.

3. **User Data Storage** (`/addDetails`):
   - Authenticated route that allows users to add their personal details.
   - Accepts user data including images and PDF files and stores them.

4. **Health Check** (`/`):
   - Simple endpoint to check if the backend server is up and running.

## Getting Started

Follow these steps to set up the project locally:

1. **Prerequisites**:
   - Install Golang on your machine: [Golang Installation Guide](https://golang.org/doc/install)
   - Install any required dependencies.

2. **Clone the Repository**:
   ```
   git clone https://github.com/whoisaditya/golang-task.git
   cd golang-task
   ```

3. **Environment Variables**:
   - Create a `.env` file based on the `.env.example` file and set your environment variables, including database credentials, JWT secret, etc.

4. **Database Setup**:
   - Set up the required database (e.g., MySQL, PostgreSQL) and update the database credentials in the `.env` file.

5. **Run the Application**:
   ```
   go run main.go
   ```

6. **Testing the Endpoints**:
   - You can now test the API endpoints using tools like `curl`, `Postman`, or `Insomnia`.
   - The base URL for the API would be `http://localhost:<PORT>/`, where `<PORT>` is the port number specified in your `.env` file.

## API Endpoints

1. **Health Check**:
   - `GET /`
   - This endpoint is used to verify if the server is up and running.

2. **User Registration**:
   - `POST /signup`
   - Request Body: `{ "name": "John Doe", "email": "john@example.com", "password": "your_password", "phone": "1234567890" }`
   - Creates a new user account and stores the provided user details in the database.

3. **User Authentication**:
   - `POST /login`
   - Request Body: `{ "email": "john@example.com", "password": "your_password" }`
   - Authenticates the user using the provided email and password.
   - Returns a JWT token upon successful authentication.

4. **User Data Storage**:
   - `POST /addDetails`
   - Request Headers: `{ "Authorization": "Bearer YOUR_JWT_TOKEN" }`
   - Request Body: Accepts various user details, including images and PDF files.
   - Stores the provided user data for the authenticated user.

## Security Considerations

- The API uses JWT for authentication, ensuring that only authenticated users can access the `/addDetails` route.
- Passwords are securely hashed before storing them in the database to protect user information.

## Contributing

If you would like to contribute to this project, you can follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them with clear commit messages.
4. Push your changes to your forked repository.
5. Submit a pull request to the `main` branch of this repository.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Documentation
Refer to the [postman documentation](https://documenter.getpostman.com/view/16151723/2s93m7V1VV)

<p align="center">
	With :heart: by <a href="https://github.com/whoisaditya" target="_blank">Aditya Mitra</a>
</p>

