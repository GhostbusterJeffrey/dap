# Dap

Dap is a web-based Platform as a Service (PaaS) that allows users to create and manage their own demo APIs effortlessly. It is designed to facilitate rapid application development by providing demo APIs that help visualize app functionality before building the complete backend logic.

[Live Service](https://dap.uranium.work)

## Features

- **User Authentication**: Secure sign-up and login through Google OAuth.
- **Project Management**: Create and manage multiple projects, each with its own API routes.
- **API Route Creation**: Add customizable API routes with defined paths, methods, and responses.
- **Demo Environment**: Quickly scaffold APIs for mobile app or website development.
- **Cloud Stoarage**: Store user and project data online, privately and securely.

## Technologies Used

- **Go**: The primary programming language for the backend.
- **Fiber**: A fast and lightweight web framework for building APIs in Go.
- **MongoDB**: A NoSQL database for data storage.
- **JWT**: JSON Web Tokens for user authentication.

## Installation
We provide a live instance of Dap [here](https://dap.uranium.work), however you can self host the API if desired

### Prerequisites

- Go 1.22.3 or higher
- MongoDB
- Node.js (if you're using a frontend)

### Clone the Repository

```bash
git clone https://github.com/GhostbusterJeffrey/dap.git
cd dap
```

### Set Up Environment Variables
Create a .env file in the root directory and add your MongoDB URI and JWT secret:
```
MONGODB_URI=mongodb://localhost:27017
JWT_SECRET=your_jwt_secret_key
```

### Install Dependencies
Run the following command to install the necessary Go packages:
```bash
go mod tidy
```

### Run the Application
Start the server using:
```bash
go run main.go
```

The server will start on ```localhost:8080``` by default.

## Contributing
Contributions are welcome! Please follow these steps to contribute:

- Fork the repository.
- Create a new branch (git checkout -b feature-branch).
- Make your changes.
- Commit your changes (git commit -m 'Add some feature').
- Push to the branch (git push origin feature-branch).
- Open a Pull Request.
License
This project is licensed under the MIT License - see the LICENSE file for details.

Contact
For questions or suggestions, feel free to reach out:
- Email: jeffrey@uranium.work
- GitHub: GhostbusterJeffrey