# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Vivek Siddhartha <siddelhi1@gmail.com>"

# Set the Current Working Directory inside the container


# Declare required environment variables
#ENV GOPATH /go
#RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH


# Get the required Go packages
RUN go get -u github.com/gorilla/securecookie
RUN go get -u github.com/gorilla/sessions
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/mediocregopher/radix.v2/pool
RUN go get -u github.com/gorilla/context
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/mediocregopher/radix.v2





# Copy go mod and sum files (Option)
#COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
#RUN go mod download 

# Copy the source from the current directory to the Working Directory inside the container
#COPY . .


# Specify the entrypoint


# Expose port 8080 to the outside world
#EXPOSE 8080

# Command to run the executable
RUN pwd
# Build the Go app
RUN go get -d -v ./...
RUN go install -v ./...


