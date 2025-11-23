# Basic Go API blueprint

This is a basic starter project template for building RESTFul API's in Go. Includes a simple project structure and Docker support.

## Usage

    # build and run without docker
    go build -o main.out main.go
    ./main.out

    # run with hot reload
    go run main.go

    # with docker
    docker build -t blueprint .
    docker run -p 8080:8080 blueprint

    # with docker-compose
    docker-compose up --build
    docker-compose down