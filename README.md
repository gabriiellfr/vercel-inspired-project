# Vercel-Inspired Project

This project replicates some core functionalities of Vercel using Go and Ansible for learning purposes.

### Project Structure

-   `app`: Contains the Go application, including handlers and templates.
-   `ansible`: Holds Ansible playbooks, roles, and variables.

### Usage

1. Clone this repository.
2. Run `go run app/main.go`.
3. Access the web interface at http://localhost:8080.
4. Provide the GitHub repository URL and project type (React or Node.js) to build and run the chosen project.

### Security Note

This example demonstrates basic functionalities. Remember to sanitize user input before using it in shell commands to prevent potential security vulnerabilities.

### Further Enhancements

-   Implement proper logging and error handling.
-   Integrate environment variables for sensitive information.
-   Explore automated image building and pushing.

This project is intended for learning purposes. Adapt and secure it appropriately for production use.
