# PortfolioAPI
The Portfolio API is a backend service designed to serve data about GitHub repositories and provide health checks for monitoring. This API supports the following routes:

* `/` - Serves static files from the `./site` directory.
* `/projects` - Fetches and returns pinned GitHub repositories.
* `/health` - Provides a simple health check endpoint.

## Project Overview
This API is built using Go and integrates with the GitHub GraphQL API to retrieve information about pinned repositories. The server also provides a health check endpoint to ensure it is operational.

## Features
* **Static File Serving:** Serves static files from the ./site directory.
* **GitHub Integration:** Fetches pinned repositories from GitHub.
* **Health Check:** Simple endpoint to check if the server is running.

## Setup and Installation

### Prerequisites

* Go installed on your machine.
* GitHub Personal Access Token with repo scope for GitHub API access.

### Steps
1. Clone the Repository:
```bash
git clone https://github.com/cameronMcConnell/PortfolioAPI.git
```

2. Navigate to the Project Directory:
```bash
cd PortfolioAPI
```

3. Install Dependencies:
```bash
go mod download
```

4. Create a .env File:
Create a file named .env in the root directory and add your GitHub API access token and server address:
```env
GITHUB_API_ACCESS_KEY=your_github_access_token
SERVER_ADDRESS=:8080
```

5. Build and Run the Server:
```bash
go build -o main .
./main
```

6. Access the API:

* **Static Files:** Access static files at `http://localhost:8080/`.
* **GitHub Projects:** Access GitHub projects at `http://localhost:8080/projects`.
* **Health Check:** Check server health at `http://localhost:8080/health`.

## Deployment
This API is designed to be deployed on AWS EC2 with the following setup:

* **EC2 Instance:** Hosts the Go application.
* **Load Balancer:** Distributes incoming traffic to the EC2 instance.
* **Route 53:** Manages DNS and directs traffic to the load balancer.
* **HTTPS Configuration:** Ensures secure communication between clients and the server.

## Contributing
Feel free to contribute by submitting issues or pull requests. Your feedback and contributions are welcome!

## License
This project is licensed under the MIT License - see the LICENSE file for details.


