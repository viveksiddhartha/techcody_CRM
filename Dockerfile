FROM golang:latest
MAINTAINER Kamesh Balasubramanian kamesh@kamesh.com

# Declare required environment variables
ENV GOPATH=/go

# Get the required Go packages
RUN go get -u github.com/gorilla/securecookie
RUN go get -u github.com/gorilla/sessions
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/mediocregopher/radix.v2/pool
RUN go get -u github.com/gorilla/context
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/mediocregopher/radix.v2


# Transpile and install the client-side application code
#RUN go get -v ./..


# Build and install the server-side application code
WORKDIR /go/src/svcrm
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...


# Specify the entrypoint
#ENTRYPOINT /go/src/svcrm/main

# Expose port 8080 of the container
#EXPOSE 8080

CMD [ "svcrm","run" ]