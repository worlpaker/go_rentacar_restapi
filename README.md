# Go Rent A Car Rest API

This project is a simple REST API for a car rental service, implemented in Go without relying on external frameworks.

- Router: native net/http - DB: MongoDB - Env: Docker

## Features

- Manual handling of routing, middlewares, logging and helper functions.

- Swagger Documentation

- Go unit tests with mocking MongoDB

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/)

## Setup

- Clone the repository:

```sh
git clone https://github.com/worlpaker/go_rentacar_restapi.git
```

- Make sure you are in the correct directory:

```sh
cd go_rentacar_restapi
```

- Before starting the services, ensure that you set the necessary environment variables in the `config/config.go`.

- Build containers and start services:

```sh
docker-compose up --build -d
```

## Access

Backend: <http://localhost:8000/>

## API Endpoints

Information about each endpoint, including request/response formats and parameters, is available in the Swagger API documentation.

- Access Docs on API: <http://localhost:8000/api/swagger/>

- Alternatively, you can also access it manually at: `api/docs`

## Running Tests

- Run tests using the `make` command:

```sh
make test
```

- To generate coverage reports and obtain more detailed information about the tests, use the following `make` command:

```sh
make cover
```
