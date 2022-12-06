# Start from golang base image
FROM golang:alpine3.16 as builder

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY .env .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the Go app
RUN go build -o main

# stage 2: copy only the application binary file and necessary files to the alpine container
FROM alpine:3.12
RUN apk --update add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .

# run the service on container startup.
CMD ["/app/main"]