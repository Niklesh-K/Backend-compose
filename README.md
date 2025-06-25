#  Backend-compose – Multi-Container Microservices with Docker Compose

This project demonstrates a simple microservices architecture using Docker and Docker Compose.

---

## Project Structure

├── docker-compose.yml # Defines and runs multi-container Docker applications
├── README.md # Project documentation
├── nginx/ # NGINX reverse proxy configuration
│ └── nginx.conf # NGINX config file
├── service_1/ # Go backend service
│ ├── Dockerfile # Dockerfile for Go service
│ ├── go.mod # Go module file
│ ├── go.sum # Go dependencies
│ └── main.go # Go source code
├── service_2/ # Python (Flask) backend service
│ ├── Dockerfile # Dockerfile for Flask service
│ ├── app.py # Flask source code
│ └── requirements.txt # Python dependencies


This structure helps explain the layout of your microservices, the role of each folder, and how your docker-compose.yml ties them together.

---

> Make sure [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/) are installed on your system.

1. Clone the repo:

```bash
git clone https://github.com/Niklesh-K/Backend-compose.git

# Build and run the containers:

  docker-compose up --build

Docker Compose will:

Build the Go service (service_1)

Build the Flask service (service_2)

Launch an NGINX reverse proxy

Create an internal Docker network


Once all services are up:

Go service will run at: http://localhost:5000

Python (Flask) service will run at: http://localhost:8000

NGINX will reverse proxy and serve: http://localhost (based on config)